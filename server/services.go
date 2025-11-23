package server

import "github.com/CrazyThursdayV50/Socrates/internal/services"

type Services struct {
	chatService services.Chat
}

func (s *Server) initServices() {
	s.services.chatService = s.impls.geminiChatter
}
