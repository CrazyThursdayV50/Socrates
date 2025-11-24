package client

import (
	"os"
	"testing"
	"time"

	"github.com/CrazyThursdayV50/Socrates/proto/chatws"
	"github.com/CrazyThursdayV50/pkgo/log/sugar"
)

func TestClient(t *testing.T) {
	var cfg Config
	cfg.URL = "ws://localhost:50505/ws"
	cfg.ReadTimeout = time.Minute
	cfg.WriteTimeout = time.Minute

	logger := sugar.New(sugar.DefaultConfig())

	client := New(logger, &cfg)

	client.HandleAnswer(func(e *chatws.Event[*chatws.AnswerData], err error) (int, []byte) {
		if err != nil {
			logger.Errorf("handler answer receive error: %v", err)
			return 1, nil
		}

		logger.Infof("answer: %v", e.Data.Answer)
		return 1, nil
	})

	err := client.Run(t.Context())
	if err != nil {
		t.Fatalf("run client failed: %v", err)
	}

	client.SetModel("gemini-2.5-flash")
	client.SetToken(os.Getenv("GEMINI_API_KEY"))
	client.Chat("我是Alex，你是谁啊？")
	time.Sleep(time.Hour)
}
