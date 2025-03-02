package web

import (
	"github.com/gin-gonic/gin"
	"github.com/snple/beacon/http/util"
	"github.com/snple/beacon/http/util/shiftime"
	"github.com/snple/beacon/pb"
	"github.com/snple/beacon/pb/cores"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SourceService struct {
	ws *WebService
}

func newSourceService(ws *WebService) *SourceService {
	return &SourceService{
		ws: ws,
	}
}

func (s *SourceService) register(router gin.IRouter) {
	group := router.Group("/source")

	group.Use(s.ws.GetAuth().MiddlewareFunc())

	group.GET("/", s.list)
	group.GET("/:id", s.get)
	group.POST("/", s.post)
	group.PATCH("/:id", s.patch)
	group.PATCH("/:id/status", s.status)
	group.DELETE("/:id", s.delete)
}

func (s *SourceService) list(ctx *gin.Context) {
	var params struct {
		util.Page `form:",inline"`
		NodeId    string `form:"node_id"`
		Tags      string `form:"tags"`
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

	request := &cores.SourceListRequest{
		Page:   page,
		NodeId: params.NodeId,
		Tags:   params.Tags,
	}

	reply, err := s.ws.Core().GetSource().List(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	items := reply.GetSource()

	shiftime.Sources(items)

	ctx.JSON(util.Success(gin.H{
		"items": items,
		"total": reply.GetCount(),
	}))
}

func (s *SourceService) get(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetSource().View(ctx, request)
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

	shiftime.Source(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *SourceService) post(ctx *gin.Context) {
	var params pb.Source

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply, err := s.ws.Core().GetSource().Create(ctx, &params)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	shiftime.Source(reply)

	ctx.JSON(util.Success(gin.H{
		"item": reply,
	}))
}

func (s *SourceService) patch(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetSource().View(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	var params pb.Source

	if err := ctx.Bind(&params); err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	reply.Name = params.Name
	reply.Desc = params.Desc
	reply.Tags = params.Tags
	reply.Source = params.Source
	reply.Params = params.Params
	reply.Config = params.Config
	reply.Status = params.Status

	reply2, err := s.ws.Core().GetSource().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}

func (s *SourceService) delete(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	_, err := s.ws.Core().GetSource().Delete(ctx, request)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": ctx.Param("id")}))
}

func (s *SourceService) status(ctx *gin.Context) {
	request := &pb.Id{Id: ctx.Param("id")}

	reply, err := s.ws.Core().GetSource().View(ctx, request)
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

	reply2, err := s.ws.Core().GetSource().Update(ctx, reply)
	if err != nil {
		ctx.JSON(util.Error(400, err.Error()))
		return
	}

	ctx.JSON(util.Success(gin.H{"id": reply2.GetId()}))
}
