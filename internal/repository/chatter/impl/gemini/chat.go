package gemini

import (
	"context"
	"errors"

	"github.com/CrazyThursdayV50/pkgo/file"
	"google.golang.org/genai"
)

func (c *Chatter) Chat(ctx context.Context, question string) (string, error) {
	if c.ai == nil {
		return "", errors.New("token not set")
	}

	var cfg = genai.GenerateContentConfig{ThinkingConfig: c.thinkingConfig}
	if c.system != "" {
		cfg.SystemInstruction = genai.NewContentFromText(c.system, genai.RoleModel)
	}

	c.logger.Debugf("system: %s", c.system)
	c.logger.Debugf("user: %s", question)
	resp, err := c.ai.Models.GenerateContent(ctx, c.model, genai.Text(question), &cfg)
	if err != nil {
		return "", err
	}

	answer := resp.Text()
	c.logger.Debugf("answer: %s", answer)
	return resp.Text(), nil
}

func (c *Chatter) SetModel(model string) {
	c.model = model
}

func (c *Chatter) SetToken(ctx context.Context, token string) error {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{APIKey: token})
	if err != nil {
		return err
	}
	c.ai = client
	return nil
}

func (c *Chatter) LoadSystem() error {
	system, err := file.ReadFileToString(c.cfg.SystemFilePath)
	if err != nil {
		return err
	}

	c.logger.Debugf("load system: %s", system)
	c.system = system
	return nil
}

func (c *Chatter) GetSystem() string { return c.system }
