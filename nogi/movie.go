package nogi

import (
	"context"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetMovies(ctx context.Context, offset, limit int) ([]*model.Movie, error) {
	panic("not implemented") // TODO: Implement
}

func (c *client) GetAllMovies(ctx context.Context) ([]*model.Movie, error) {
	panic("not implemented") // TODO: Implement
}
