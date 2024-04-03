package web

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kokomi/http/util"
	"github.com/snple/kokomi/http/util/shiftime"
	"github.com/snple/kokomi/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeviceService struct {
	ws *WebService
}

func newDeviceService(ws *WebService) *DeviceService {
	return &DeviceService{
		ws: ws,
	}
}

func (s *DeviceService) register(router gin.IRouter) {
	group := router.Group("/device")

	group.Use(s.ws.GetAuth().MiddlewareFunc())

	group.GET("/id", s.get)
	group.PATCH("/id", s.patch)
	group.PATCH("/id/status", s.status)
}

func (s *DeviceService) get(ctx *gin.Context) {
	reply, err := s.ws.Edge().GetDevice().View(ctx, &pb.MyEmpty{})
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

func (s *DeviceService) patch(ctx *gin.Context) {
	reply, err := s.ws.Edge().GetDevice().View(ctx, &pb.MyEmpty{})
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params pb.Device

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Name = params.Name
	reply.Desc = params.Desc
	reply.Tags = params.Tags
	reply.Type = params.Type
	reply.Secret = params.Secret
	reply.Location = params.Location
	reply.Config = params.Config
	reply.Status = params.Status

	reply2, err := s.ws.Edge().GetDevice().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}

func (s *DeviceService) status(ctx *gin.Context) {
	reply, err := s.ws.Edge().GetDevice().View(ctx, &pb.MyEmpty{})
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params struct {
		Status int32 `json:"status"`
	}

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Status = params.Status

	reply2, err := s.ws.Edge().GetDevice().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}
