package chatws

import "github.com/CrazyThursdayV50/pkgo/json"

const (
	EVENET_ANSWER = "answer"
)

type Event[T any] struct {
	ID    int64  `json:"id"`
	Event string `json:"event"`
	Data  T
}

func (e *Event[T]) MarshalBinary() ([]byte, error) {
	return json.JSON().Marshal(e)
}

type ResultData struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type AnswerData struct {
	Answer string `json:"answer"`
}

func EventAnswer(id int64, answer string) *Event[*AnswerData] {
	return &Event[*AnswerData]{
		ID:    id,
		Event: EVENET_ANSWER,
		Data: &AnswerData{
			Answer: answer,
		},
	}
}

func EventResultOK(id int64, event string) *Event[*ResultData] {
	return &Event[*ResultData]{
		ID:    id,
		Event: event,
		Data: &ResultData{
			OK:      true,
			Error:   "",
			Message: "",
		},
	}
}

func EventResultFail(id int64, event string, err, message string) *Event[*ResultData] {
	return &Event[*ResultData]{
		ID:    id,
		Event: event,
		Data: &ResultData{
			OK:      false,
			Error:   err,
			Message: message,
		},
	}
}
