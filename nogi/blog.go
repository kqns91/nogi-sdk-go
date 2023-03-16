package nogi

import (
	"context"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetBlogs(ctx context.Context, offset, limit int) ([]*model.Blog, error) {
	panic("not implemented") // TODO: Implement
}

func (c *client) GetAllBlogs(ctx context.Context) ([]*model.Blog, error) {
	panic("not implemented") // TODO: Implement
}
