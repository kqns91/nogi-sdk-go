package nogi

import (
	"context"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetLives(ctx context.Context, offset, limit int) ([]*model.Live, error) {
	panic("not implemented") // TODO: Implement
}

func (c *client) GetAllLives(ctx context.Context) ([]*model.Live, error) {
	panic("not implemented") // TODO: Implement
}
