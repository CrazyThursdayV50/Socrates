package gemini

import (
	"net/http"

	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter"
	"github.com/CrazyThursdayV50/Socrates/internal/services/wshandler"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/trace"
	"github.com/CrazyThursdayV50/pkgo/websocket/server"
	"github.com/gin-gonic/gin"
)

type Service struct {
	logger   log.Logger
	wsserver *server.Server
	chatter  chatter.Repository
}

func New(logger log.Logger, tracer trace.Tracer, chatter chatter.Repository) *Service {
	server := server.New(
		server.WithLogger(logger),
		server.WithTracer(tracer),
		server.WithHandler(wshandler.NewChatterHandler(logger, chatter)),
	)

	return &Service{
		logger:   logger,
		wsserver: server,
		chatter:  chatter,
	}
}

func (c *Service) Listen(ctx *gin.Context) {
	c.logger.Debugf("ws incoming new conn")
	err := c.wsserver.Run(ctx, ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, struct{}{})
		return
	}
	ctx.JSON(http.StatusOK, struct{}{})
}

func (c *Service) ReloadSystem(ctx *gin.Context) {
	err := c.chatter.LoadSystem()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "reload system failed: %v", err)
		return
	}

	ctx.String(http.StatusOK, c.chatter.GetSystem())
}
