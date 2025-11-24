package gemini

import (
	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter"
	"github.com/CrazyThursdayV50/pkgo/log"
	"google.golang.org/genai"
)

type Chatter struct {
	cfg            *chatter.Config
	logger         log.Logger
	thinkingConfig *genai.ThinkingConfig

	ai     *genai.Client
	system string
	model  string
}

func New(cfg *chatter.Config, logger log.Logger) *Chatter {
	var c Chatter
	c.logger = logger
	c.cfg = cfg
	c.thinkingConfig = &genai.ThinkingConfig{IncludeThoughts: false}
	return &c
}
