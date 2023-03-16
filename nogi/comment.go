package nogi

import (
	"context"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetAllComments(ctx context.Context, blogID string) ([]*model.Comment, error) {
	panic("not implemented") // TODO: Implement
}

func (c *client) GetComments(ctx context.Context, blogID string, offset, limit int) ([]*model.Comment, error) {
	panic("not implemented") // TODO: Implement
}
