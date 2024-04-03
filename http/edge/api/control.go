package api

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/kokomi/http/util"
	"github.com/snple/kokomi/pb"
)

type ControlService struct {
	as *ApiService
}

func newControlService(as *ApiService) *ControlService {
	return &ControlService{
		as: as,
	}
}

func (s *ControlService) register(router gin.IRouter) {
	group := router.Group("/control")

	group.GET("/tag/:id/value", s.getTagValue)
	group.PATCH("/tag/:id/value", s.setTagValue)
}

func (s *ControlService) getTagValue(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.as.Edge().GetControl().GetTagValue(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *ControlService) setTagValue(ctx *gin.Context) {
	var request pb.TagValue

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	request.Id = ctx.Param("id")

	reply, err := s.as.Edge().GetControl().SetTagValue(ctx, &request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}
