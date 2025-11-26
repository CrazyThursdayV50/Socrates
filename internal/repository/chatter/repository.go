package chatter

import "context"

type Repository interface {
	SetModel(model string)
	SetToken(_ context.Context, token string) error
	LoadSystem() error
	GetSystem() string
	Chat(ctx context.Context, question string) (string, error)
}
