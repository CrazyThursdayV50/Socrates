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
	return func(ctx context.Context, messageType int, data []byte, err error) (frame int, bytes []byte, e error) {
		frame = 2

		logger.Infof("Recv: %s", data)
		defer func() {
			if bytes != nil {
				logger.Debugf("Send: %s", bytes)
			}
		}()

		id := chatws.GetID(data)
		name := chatws.GetName(data)

		switch name {
		case chatws.ACTION_QUESTION:
			user := gjson.GetBytes(data, "data.user").String()
			answer, err := chat.Chat(ctx, user)

			if err != nil {
				bytes, _ = chatws.EventResultFail(id, chatws.EVENET_ANSWER, "chat failed", err.Error()).MarshalBinary()
			} else {
				bytes, _ = chatws.EventAnswer(id, answer).MarshalBinary()
			}

			return

		case chatws.ACTION_SET_CONFIG:
			result := gjson.GetBytes(data, "data.token")
			if result.Exists() {
				err := chat.SetToken(ctx, result.String())
				if err != nil {
					bytes, _ = chatws.EventResultFail(id, name, "set token failed", err.Error()).MarshalBinary()
					return
				}
			}

			result = gjson.GetBytes(data, "data.model")
			if result.Exists() {
				chat.SetModel(result.String())
			}

			bytes, _ = json.JSON().Marshal(chatws.EventResultOK(id, name))
			return

		default:
			logger.Warnf("unexpected name: %v", name)
		}

		return
	}
}
