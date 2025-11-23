package server

import "github.com/CrazyThursdayV50/Socrates/internal/repository/chatter/impl/gemini"

type Impls struct {
	geminiChatter *gemini.Chatter
}

func (s *Server) initImpls() {
	s.impls.geminiChatter = gemini.New(s.logger, s.tracerCreator.NewTracer("GeminiChatter"))
}
