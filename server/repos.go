package server

import (
	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter"
)

type Repos struct {
	geminiChatter chatter.Repository
}

func (s *Server) initRepos() {
	s.repos.geminiChatter = s.impls.geminiChatter
}
