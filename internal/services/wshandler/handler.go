package wshandler

import (
	"context"

	"github.com/CrazyThursdayV50/Socrates/internal/repository/chatter"
	"github.com/CrazyThursdayV50/Socrates/proto/chatws"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/tidwall/gjson"
)

func NewChatterHandler(logger log.Logger, chat chatter.Repository) func(ctx context.Context, messageType int, data []byte, err error) (int, []byte, error) {
	return func(ctx context.Context, messageType int, data []byte, err error) (int, []byte, error) {
		logger.Infof("Recv: %s", data)

		id := chatws.GetID(data)
		name := chatws.GetName(data)

		switch name {
		case chatws.ACTION_QUESTION:
			user := gjson.GetBytes(data, "data.user").String()
			answer, err := chat.Chat(ctx, user)

			if err != nil {
				data, _ = chatws.EventResultFail(id, chatws.EVENET_ANSWER, "chat failed", err.Error()).MarshalBinary()
			} else {
				data, _ = chatws.EventAnswer(id, answer).MarshalBinary()
			}

			return 2, data, nil

		case chatws.ACTION_SET_CONFIG:
			result := gjson.GetBytes(data, "data.token")
			if result.Exists() {
				err := chat.SetToken(ctx, result.String())
				if err != nil {
					data, _ = chatws.EventResultFail(id, name, "set token failed", err.Error()).MarshalBinary()
					return 2, data, nil
				}
			}

			result = gjson.GetBytes(data, "data.system")
			if result.Exists() {
				chat.SetSystem(result.String())
			}

			result = gjson.GetBytes(data, "data.model")
			if result.Exists() {
				chat.SetModel(result.String())
			}

			data, _ = json.JSON().Marshal(chatws.EventResultOK(id, name))
			return 2, data, nil

		default:
			logger.Warnf("unexpected name: %v", name)
		}

		return 1, nil, nil
	}
}
