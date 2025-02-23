package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/beacon/http/util"
	"github.com/snple/beacon/http/util/shiftime"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceService struct {
	as *ApiService
}

func newDeviceService(as *ApiService) *DeviceService {
	return &DeviceService{
		as: as,
	}
}

func (s *DeviceService) register(router gin.IRouter) {
	group := router.Group("/device")

	group.GET("/", s.list)

	group.GET("/:id", s.getById)

	group.GET("/name/:name", s.getByName)
	group.POST("/name", s.getByNames)

	group.PATCH("/link", s.link)
}

func (s *DeviceService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		Tags      string `form:"tags"`
		Type      string `form:"type"`
	}

	if err := ctx.Bind(&params); err != nil {
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

	request := &cores.DeviceListRequest{
		Page: page,
		Tags: params.Tags,
		Type: params.Type,
	}

	reply, err := s.as.Core().GetDevice().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{
		"items": reply.GetDevice(),
		"total": reply.GetCount(),
	}))
}

func (s *DeviceService) getById(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Core().GetDevice().View(ctx, request)
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

func (s *DeviceService) getByName(ctx *gin.Context) {
	name := ctx.Param("name")

	reply, err := s.as.Core().GetDevice().Name(ctx, &pb.Name{Name: name})
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

	shiftime.Device(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *DeviceService) getByNames(ctx *gin.Context) {
	var params struct {
		Name []string `json:"name"`
	}
	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ret := make([]*pb.Device, 0, len(params.Name))

	for _, name := range params.Name {
		reply, err := s.as.Core().GetDevice().Name(ctx,
			&pb.Name{Name: name})
		if err != nil {
			if code, ok := status.FromError(err); ok {
				if code.Code() == codes.NotFound {
					continue
				}
			}

			ctx.JSON(util.Error(400, err.Error()))
			return
		}

		ret = append(ret, reply)
	}

	ctx.JSON(util.Success(ret))
}

func (s *DeviceService) link(ctx *gin.Context) {
	var params struct {
		Id     string `json:"id"`
		Status int    `json:"status"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.as.Core().GetDevice().Link(ctx,
		&cores.DeviceLinkRequest{Id: params.Id, Status: int32(params.Status)})
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

	ctx.JSON(util.Success(reply))
}
