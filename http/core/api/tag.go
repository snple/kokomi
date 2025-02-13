package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kokomi/http/util"
	"github.com/snple/kokomi/http/util/shiftime"
	"github.com/snple/kokomi/pb"
	"github.com/snple/kokomi/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TagService struct {
	as *ApiService
}

func newTagService(as *ApiService) *TagService {
	return &TagService{
		as: as,
	}
}

func (s *TagService) register(router gin.IRouter) {
	group := router.Group("/tag")

	group.GET("/", s.list)

	group.GET("/:id", s.getById)
	group.GET("/:id/value", s.getValueById)
	group.PATCH("/:id/value", s.setValueById)

	group.GET("/name/:name", s.getByName)
	group.POST("/name", s.getByNames)

	group.POST("/get_value", s.getValueByNames)
	group.PATCH("/set_value", s.setValueByNames)

	group.POST("/get_write", s.getWriteByNames)
	group.PATCH("/set_write", s.setWriteByNames)
}

func (s *TagService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		DeviceId  string `form:"device_id"`
		Name      string `form:"name"`
		Tags      string `form:"tags"`
		Type      string `form:"type"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	source, err := s.as.Core().GetSource().Name(ctx,
		&cores.SourceNameRequest{DeviceId: params.DeviceId, Name: params.Name})
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

	page := &pb.Page{
		Limit:   params.Limit,
		Offset:  params.Offset,
		Search:  params.Search,
		OrderBy: params.OrderBy,
		Sort:    pb.Page_ASC,
	}

	if params.Sort > 0 {
		page.Sort = pb.Page_DESC
	}

	request := &cores.TagListRequest{
		Page:     page,
		DeviceId: params.DeviceId,
		SourceId: source.Id,
		Tags:     params.Tags,
		Type:     params.Type,
	}

	reply, err := s.as.Core().GetTag().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetTag()

	shiftime.Tags(items)

	ctx.JSON(util.Success(gin.H{
		"source": source,
		"items":  items,
		"total":  reply.GetCount(),
	}))
}

func (s *TagService) getById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Core().GetTag().View(ctx, request)
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

	shiftime.Tag(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *TagService) getValueById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Core().GetTag().GetValue(ctx, request)
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

	shiftime.TagValue(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *TagService) setValueById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	var params struct {
		Value string `json:"value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Core().GetTag().SetValue(ctx,
		&pb.TagValue{Id: request.Id, Value: params.Value})
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

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *TagService) getByName(ctx *gin.Context) {
	name := ctx.Param("name")

	var params struct {
		DeviceId string `form:"device_id"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Core().GetTag().Name(ctx,
		&cores.TagNameRequest{DeviceId: params.DeviceId, Name: name})
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

	shiftime.Tag(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *TagService) getByNames(ctx *gin.Context) {
	var params struct {
		DeviceId string   `json:"device_id"`
		Name     []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*pb.Tag, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Core().GetTag().Name(ctx,
			&cores.TagNameRequest{DeviceId: params.DeviceId, Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftime.Tag(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *TagService) getValueByNames(ctx *gin.Context) {
	var params struct {
		DeviceId string   `json:"device_id"`
		Name     []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*cores.TagNameValue, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Core().GetTag().GetValueByName(ctx,
			&cores.TagGetValueByNameRequest{DeviceId: params.DeviceId, Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftTimeForTagNameValue(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *TagService) setValueByNames(ctx *gin.Context) {
	var params struct {
		DeviceId  string            `json:"device_id"`
		NameValue map[string]string `json:"name_value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	errors := make(map[string]string)

	for name, value := range params.NameValue {
		_, err := s.as.Core().GetTag().SetValueByName(ctx,
			&cores.TagNameValue{DeviceId: params.DeviceId, Name: name, Value: value})
		if err != nil {
			errors[name] = err.Error()
		}
	}

	if len(errors) > 0 {
		ctx.JSON(util.Success(gin.H{
			"ok":     false,
			"errors": errors,
		}))

		return
	}

	ctx.JSON(util.Success(gin.H{
		"ok": true,
	}))
}

func (s *TagService) getWriteByNames(ctx *gin.Context) {
	var params struct {
		DeviceId string   `json:"device_id"`
		Name     []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*cores.TagNameValue, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Core().GetTag().GetWriteByName(ctx,
			&cores.TagGetValueByNameRequest{DeviceId: params.DeviceId, Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		shiftTimeForTagNameValue(reply)

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *TagService) setWriteByNames(ctx *gin.Context) {
	var params struct {
		DeviceId  string            `json:"device_id"`
		NameValue map[string]string `json:"name_value"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	errors := make(map[string]string)

	for name, value := range params.NameValue {
		_, err := s.as.Core().GetTag().SetWriteByName(ctx,
			&cores.TagNameValue{DeviceId: params.DeviceId, Name: name, Value: value})
		if err != nil {
			errors[name] = err.Error()
		}
	}

	if len(errors) > 0 {
		ctx.JSON(util.Success(gin.H{
			"ok":     false,
			"errors": errors,
		}))

		return
	}

	ctx.JSON(util.Success(gin.H{
		"ok": true,
	}))
}

func shiftTimeForTagNameValue(item *cores.TagNameValue) {
	if item != nil {
		item.Updated = item.Updated / 1000
	}
}

func shiftTimeForTagNameValues(items []*cores.TagNameValue) {
	for _, item := range items {
		shiftTimeForTagNameValue(item)
	}
}
