package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/beacon/http/util"
	"github.com/snple/beacon/pb"
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

	group.GET("/", s.view)
}

func (s *DeviceService) view(ctx *gin.Context) {
	reply, err := s.as.Edge().GetDevice().View(ctx, &pb.MyEmpty{})
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
