package gemini

import (
	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter/impl/gemini"
	"github.com/CrazyThursdayV50/Socrates/internal/services"
)

type Chatter = gemini.Chatter

var _ services.Chat = (*Chatter)(nil)

var New = gemini.New
