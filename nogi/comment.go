package nogi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
)

func (c *client) GetComments(ctx context.Context, blogID string, offset, limit int) ([]*model.Comment, error) {
	query := url.Values{}
	query.Set("kiji", blogID)
	query.Set("st", strconv.Itoa(offset))
	query.Set("rw", strconv.Itoa(limit))

	urlStr := fmt.Sprintf("%s://%s/%s/comment?%s", c.httpProtocol, baseURL, apiBasePath, query.Encode())

	data, err := c.httpClient.send(ctx, http.MethodGet, urlStr, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send: %w", err)
	}

	res := &model.Comments{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	html := regexp.MustCompile("<[^>]*>")

	for _, d := range res.Data {
		d.Body = html.ReplaceAllString(d.Body, "")
	}

	return res.Data, nil
}

func (c *client) GetAllComments(ctx context.Context, blogID string) ([]*model.Comment, error) {
	result := []*model.Comment{}

	st := 0
	rw := 10
	query := url.Values{}
	query.Set("kiji", blogID)
	query.Set("rw", strconv.Itoa(rw))

	for {
		query.Set("st", strconv.Itoa(st))

		urlStr := fmt.Sprintf("%s://%s/%s/comment?%s", c.httpProtocol, baseURL, apiBasePath, query.Encode())

		data, err := c.httpClient.send(ctx, http.MethodGet, urlStr, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to send: %w", err)
		}

		res := &model.Comments{}
		err = json.Unmarshal(data, res)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal: %w", err)
		}

		html := regexp.MustCompile("<[^>]*>")

		for _, d := range res.Data {
			d.Body = html.ReplaceAllString(d.Body, "")
		}

		result = append(result, res.Data...)

		st += rw

		count, err := strconv.Atoi(res.Count)
		if err != nil {
			return nil, fmt.Errorf("failed to convert count: %q", res.Count)
		}

		if count <= len(result) {
			break
		}
	}

	return result, nil
}
