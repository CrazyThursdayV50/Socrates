package server

import (
	"github.com/CrazyThursdayV50/Socrates/internal/services"
	"github.com/CrazyThursdayV50/Socrates/internal/services/impl/gemini"
)

type Services struct {
	chatService services.Chat
}

func (s *Server) initServices() {
	s.services.chatService = gemini.New(
		s.logger,
		s.tracerCreator.NewTracer("ServiceGemini"),
		s.impls.geminiChatter,
	)
}
