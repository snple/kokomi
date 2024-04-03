package api

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snple/kokomi/consts"
	"github.com/snple/kokomi/core/model"
	"github.com/snple/kokomi/db"
	"github.com/snple/kokomi/http/util"
	"github.com/snple/kokomi/http/util/binding"
	modutil "github.com/snple/kokomi/util"
	"github.com/snple/kokomi/util/datatype"
	"github.com/snple/kokomi/util/flux"
	"github.com/snple/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DataService struct {
	as *ApiService
}

func newDataService(as *ApiService) *DataService {
	return &DataService{
		as: as,
	}
}

func (s *DataService) register(router gin.IRouter) {
	group := router.Group("/data")

	group.POST("/compile", s.compile)
	group.POST("/query", s.query)
	group.GET("/query/:id", s.queryById)

	group.PATCH("/upload1", s.upload1)

	group.GET("/query_by_id/:id", s.queryById) // deprecated
}

func (s *DataService) compile(ctx *gin.Context) {
	var params struct {
		Flux string            `json:"flux"`
		Vars map[string]string `json:"vars"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if len(params.Flux) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid flux"))
		return
	}

	influx := s.as.Core().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	if len(params.Vars) > 0 {
		for key, val := range params.Vars {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(params.Flux)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
	}))
}

func (s *DataService) query(ctx *gin.Context) {
	var params struct {
		Flux string            `json:"flux"`
		Vars map[string]string `json:"vars"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if len(params.Flux) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid flux"))
		return
	}

	influx := s.as.Core().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
	}

	if len(params.Vars) > 0 {
		for key, val := range params.Vars {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["bucket"] = influxdb.Bucket()

	tmpl, err := template.New("").Parse(params.Flux)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	result, err := influxdb.QueryAPI().Query(ctx, buffer.String())
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("InfluxDB Query: %v", err)))
		return
	}

	rows := make([]map[string]interface{}, 0)

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = modutil.TimeFormat(t)
			}
		}

		rows = append(rows, values)
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
		"rows": rows,
	}))
}

func (s *DataService) queryById(ctx *gin.Context) {
	id := ctx.Param("id")

	if len(id) == 0 {
		ctx.JSON(util.Error(400, "Please supply valid id"))
		return
	}

	influx := s.as.Core().GetInfluxDB()
	if influx.IsNone() {
		ctx.JSON(util.Error(400, "Influxdb is not enable"))
		return
	}
	influxdb := influx.Unwrap()

	vars2 := make(map[string]string, 0)

	if ctx.Request != nil {
		for key, value := range ctx.Request.URL.Query() {
			if len(value) > 0 {
				vars2[key] = value[0]
			}
		}
	}

	vars := map[string]string{
		"start":  "-1h",
		"stop":   "now()",
		"window": "10m",
		"fn":     "mean",
		"fill":   "none",
	}

	if len(vars2) > 0 {
		for key, val := range vars2 {
			vars[key] = tryConvertTimeZone(val)
		}
	}

	vars["id"] = id
	vars["bucket"] = influxdb.Bucket()

	script := flux.TemplateBasic

	if fill, ok := vars["fill"]; ok {
		switch fill {
		case "prev":
			script = flux.TemplateFill
		case "linear":
			script = flux.TemplateLinear
		}
	}

	tmpl, err := template.New("").Parse(script)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template parse: %v", err)))
		return
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, vars)
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("Template exec: %v", err)))
		return
	}

	result, err := influxdb.QueryAPI().Query(ctx, buffer.String())
	if err != nil {
		ctx.JSON(util.Error(400, fmt.Sprintf("InfluxDB Query: %v", err)))
		return
	}

	rows := make([]map[string]interface{}, 0)

	for result.Next() {
		values := result.Record().Values()

		for key, val := range values {
			if t, ok := val.(time.Time); ok {
				values[key] = modutil.TimeFormat(t)
			}
		}

		rows = append(rows, values)
	}

	ctx.JSON(util.Success(gin.H{
		"flux": buffer.String(),
		"rows": rows,
	}))
}

func (s *DataService) upload1(ctx *gin.Context) {
	const (
		UPLOAD_TYPE_NAME    uint32 = 0
		UPLOAD_TYPE_ADDRESS uint32 = 1
	)

	var params struct {
		DeviceId  string         `json:"device_id"`
		Realtime  bool           `json:"realtime"`
		Save      bool           `json:"save"`
		Timestamp uint32         `json:"ts"`     // 可选
		Source    string         `json:"source"` // 可选，默认 ’source‘
		Type      uint32         `json:"type"`   // 可选，默认 0
		Values    map[string]any `json:"values"`
	}

	if err := ctx.MustBindWith(&params, binding.JSONUseNumber); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	device, err := s.as.Core().GetDevice().ViewByID(ctx, params.DeviceId)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				ctx.JSON(util.Error(404, err.Error()))
				return
			}
		}

		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if device.Status != consts.ON {
		ctx.JSON(util.Error(412, "Device Status != ON"))
		return
	}

	// timestamp
	timestamp := time.Now().Unix()
	if params.Timestamp > 0 {
		timestamp = int64(params.Timestamp)
	}

	// source name
	sourceName := consts.DEFAULT_SOURCE
	if params.Source != "" {
		sourceName = params.Source
	}

	// upload type
	uploadType := UPLOAD_TYPE_NAME
	if params.Type == UPLOAD_TYPE_NAME || params.Type == UPLOAD_TYPE_ADDRESS {
		uploadType = params.Type
	}

	// get and check source
	source, err := s.as.Core().GetSource().ViewFromCacheByDeviceIDAndName(ctx, params.DeviceId, sourceName)
	if err != nil {
		if code, ok := status.FromError(err); ok {
			if code.Code() == codes.NotFound {
				ctx.JSON(util.Error(404, err.Error()))
				return
			}
		}

		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	if source.Status != consts.ON {
		ctx.JSON(util.Error(412, "Source Status != ON"))
		return
	}

	// InfluxDB writer
	writer := types.None[*db.Writer]()
	option := s.as.Core().GetInfluxDB()
	if option.IsSome() {
		writer = types.Some(option.Unwrap().Writer(50))
	}

	for name, value := range params.Values {
		// get and check tag
		var tag model.Tag
		if uploadType == UPLOAD_TYPE_ADDRESS {
			tag, err = s.as.Core().GetTag().ViewBySourceIDAndAddress(ctx, source.ID, name)
		} else {
			tag, err = s.as.Core().GetTag().ViewBySourceIDAndName(ctx, source.ID, name)
		}
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					ctx.JSON(util.Error(404, err.Error()))
					return
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		if tag.Status != consts.ON {
			continue
		}

		// validation data type
		value2, err := s.as.Core().GetData().CheckDataTypeJson(&tag, value)
		if err != nil {
			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		// realtime
		if params.Realtime {
			s.as.Core().GetData().UpdateValue(&tag, value2)
		}

		// save
		if writer.IsSome() && params.Save && source.Save == consts.ON && tag.Save == consts.ON {
			if value2, ok := datatype.NsonValueToFloat64(value2); ok {

				point := s.as.Core().GetData().NewPoint(&tag, value2, timestamp)

				err = writer.Unwrap().Write(ctx, point)
				if err != nil {
					ctx.JSON(util.Error(500, fmt.Sprintf("Write: %v", err)))
					return
				}
			}
		}
	}

	if writer.IsSome() {
		err = writer.Unwrap().Flush(ctx)
		if err != nil {
			ctx.JSON(util.Error(500, fmt.Sprintf("Flush: %v", err)))
			return
		}
	}
}

func tryConvertTimeZone(before string) string {
	t, err := modutil.ParseTime(before)
	if err == nil {
		return t.UTC().Format("2006-01-02T15:04:05Z")
	}

	return before
}
