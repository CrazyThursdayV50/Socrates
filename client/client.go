package client

import (
	"context"
	"encoding"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/CrazyThursdayV50/Socrates/proto/chatws"
	"github.com/CrazyThursdayV50/pkgo/json"
	"github.com/CrazyThursdayV50/pkgo/log"
	"github.com/CrazyThursdayV50/pkgo/websocket/client"
	"github.com/gorilla/websocket"
)

type Client struct {
	id       int64
	logger   log.Logger
	wsclient *client.Client
	handlers map[string]func([]byte) (int, []byte)

	model  string
	token  string
	system string
}

func (c *Client) Send(action encoding.BinaryMarshaler) {
	data, _ := action.MarshalBinary()
	c.logger.Debugf("Send: %s", data)
	c.wsclient.Send(data)
}

func createHandler[T any](handler func(T, error) (int, []byte)) func([]byte) (int, []byte) {
	return func(b []byte) (int, []byte) {
		var t T
		err := json.JSON().Unmarshal(b, &t)
		return handler(t, err)
	}
}

func (c *Client) SetToken(token string) {
	c.token = token
	action := chatws.ActionSetToken(token)
	action.ID = atomic.AddInt64(&c.id, 1)
	c.Send(action)
}

func (c *Client) SetModel(model string) {
	c.model = model
	action := chatws.ActionSetModel(model)
	action.ID = atomic.AddInt64(&c.id, 1)
	c.Send(action)
}

func (c *Client) SetSystem(system string) {
	c.system = system
	action := chatws.ActionSetSystem(system)
	action.ID = atomic.AddInt64(&c.id, 1)
	c.Send(action)
}

func (c *Client) HandleAnswer(handler func(*chatws.Event[*chatws.AnswerData], error) (int, []byte)) {
	h := createHandler(handler)
	c.handlers[chatws.EVENET_ANSWER] = h
}

func (c *Client) Chat(question string) {
	action := chatws.ActionQuestion(question)
	action.ID = atomic.AddInt64(&c.id, 1)
	c.Send(action)
}

func New(logger log.Logger, cfg *Config) *Client {
	var c Client
	c.logger = logger

	wsclient := client.New(
		client.WithLogger(logger),
		client.WithURL(cfg.URL),
		client.WithReadTimeout(cfg.ReadTimeout),
		client.WithWriteTimeout(cfg.WriteTimeout),
		client.WithMessageHandler(func(ctx context.Context, l log.Logger, i int, b []byte, f func(error)) (int, []byte) {
			name := chatws.GetEvent(b)
			handler, ok := c.handlers[name]
			if !ok {
				logger.Warnf("%s event handler not found: %s", name, b)
				return client.BinaryMessage, nil
			}

			return handler(b)
		}),

		client.WithPingLoop(func(done <-chan struct{}, conn *websocket.Conn) {
			for {
				select {
				case <-done:
					return

				default:
					now := fmt.Appendf(nil, "%d", time.Now().UnixMilli())
					conn.WriteControl(client.PingMessage, now, time.Now().Add(time.Second*5))
				}
			}
		}),
	)

	wsclient.UpdateOptions(client.WithOnConnect(func() (int, []byte) {
		action := chatws.ActionSetToken(c.token)
		action.Data.Model = &c.model
		action.Data.System = &c.system
		action.ID = atomic.AddInt64(&c.id, 1)
		data, _ := action.MarshalBinary()
		return client.BinaryMessage, data
	}))

	c.wsclient = wsclient
	c.handlers = make(map[string]func([]byte) (int, []byte))
	return &c
}

func (c *Client) Run(ctx context.Context) error {
	return c.wsclient.Run(ctx)
}
