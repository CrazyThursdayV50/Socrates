package gemini

import (
	"github.com/CrazyThursdayV50/Socrates/internal/services/wshandler"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/trace"
	"github.com/CrazyThursdayV50/pkgo/websocket/server"
	"google.golang.org/genai"
)

type Chatter struct {
	logger         log.Logger
	wsserver       *server.Server
	ai             *genai.Client
	system         string
	model          string
	thinkingConfig *genai.ThinkingConfig
}

func New(logger log.Logger, tracer trace.Tracer) *Chatter {
	var c Chatter
	c.logger = logger
	server := server.New(
		server.WithLogger(logger),
		server.WithTracer(tracer),
		server.WithHandler(wshandler.NewChatterHandler(logger, &c)),
	)

	c.wsserver = server
	c.thinkingConfig = &genai.ThinkingConfig{IncludeThoughts: false}
	return &c
}
