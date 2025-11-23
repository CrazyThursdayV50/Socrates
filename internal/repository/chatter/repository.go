package chatter

import "context"

type Repository interface {
	SetSystem(system string)
	SetModel(model string)
	SetToken(_ context.Context, token string) error
	Chat(ctx context.Context, question string) (string, error)
}
