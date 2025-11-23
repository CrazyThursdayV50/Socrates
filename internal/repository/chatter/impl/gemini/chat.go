package gemini

import (
	"context"

	"google.golang.org/genai"
)

func (c *Chatter) Chat(ctx context.Context, question string) (string, error) {
	var cfg = genai.GenerateContentConfig{ThinkingConfig: c.thinkingConfig}
	if c.system != "" {
		cfg.SystemInstruction = genai.NewContentFromText(c.system, genai.RoleModel)
	}

	resp, err := c.ai.Models.GenerateContent(ctx, c.model, genai.Text(question), &cfg)
	if err != nil {
		return "", err
	}
	return resp.Text(), nil
}

func (c *Chatter) SetSystem(system string) {
	c.system = system
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
