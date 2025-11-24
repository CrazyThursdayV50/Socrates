package server

import (
	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter"
)

type Impls struct {
	geminiChatter chatter.Repository
}
