package nogi

import (
	"context"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetHandShakes(ctx context.Context, offset, limit int) ([]*model.HandShake, error) {
	panic("not implemented") // TODO: Implement
}

func (c *client) GetAllHandShakes(ctx context.Context) ([]*model.HandShake, error) {
	panic("not implemented") // TODO: Implement
}
