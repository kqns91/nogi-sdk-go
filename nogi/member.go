package nogi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetAllMembers(ctx context.Context) ([]*model.Member, error) {
	urlStr := fmt.Sprintf("%s://%s/%s/member", c.httpProtocol, baseURL, apiBasePath)

	data, err := c.httpClient.send(ctx, http.MethodGet, urlStr, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send: %w", err)
	}

	res := &model.Members{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	return res.Data, nil
}
