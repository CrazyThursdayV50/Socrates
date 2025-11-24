package chatws

import (
	"github.com/CrazyThursdayV50/pkgo/json"
)

const (
	ACTION_SET_CONFIG = "set_config"
	ACTION_QUESTION   = "question"
)

type Action[T any] struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Data T      `json:"data"`
}

func (a *Action[T]) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(a)
}

type SetConfigData struct {
	Token *string `json:"token,omitempty"`
	Model *string `json:"model,omitempty"`
}

type QuestionData struct {
	User string `json:"user"`
}

func ActionQuestion(question string) *Action[*QuestionData] {
	return &Action[*QuestionData]{
		Name: ACTION_QUESTION,
		Data: &QuestionData{
			User: question,
		},
	}
}

func ActionSetToken(token string) *Action[*SetConfigData] {
	return &Action[*SetConfigData]{
		Name: ACTION_SET_CONFIG,
		Data: &SetConfigData{
			Token: &token,
		},
	}
}

func ActionSetModel(model string) *Action[*SetConfigData] {
	return &Action[*SetConfigData]{
		Name: ACTION_SET_CONFIG,
		Data: &SetConfigData{
			Model: &model,
		},
	}
}
