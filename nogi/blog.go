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

func (c *client) GetBlogs(ctx context.Context, offset, limit int) ([]*model.Blog, error) {
	query := url.Values{}
	query.Set("st", strconv.Itoa(offset))
	query.Set("rw", strconv.Itoa(limit))

	urlStr := fmt.Sprintf("%s://%s/%s/blog?%s", c.httpProtocol, baseURL, apiBasePath, query.Encode())

	data, err := c.httpClient.send(ctx, http.MethodGet, urlStr, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to send: %w", err)
	}

	res := &model.Blogs{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal: %w", err)
	}

	img := regexp.MustCompile(`<img[^>]*src="(.*?)"`)
	html := regexp.MustCompile("<[^>]*>")

	for _, d := range res.Data {
		matches := img.FindAllStringSubmatch(d.Text, -1)
		for _, match := range matches {
			d.Images = append(d.Images, fmt.Sprintf("%s://%s%s", c.httpProtocol, baseURL, match[1]))
		}

		d.Text = html.ReplaceAllString(d.Text, "")
	}

	return res.Data, nil
}

func (c *client) GetAllBlogs(ctx context.Context) ([]*model.Blog, error) {
	result := []*model.Blog{}

	st := 0
	rw := 32
	query := url.Values{}
	query.Set("rw", strconv.Itoa(rw))

	for {
		query.Set("st", strconv.Itoa(st))

		urlStr := fmt.Sprintf("%s://%s/%s/blog?%s", c.httpProtocol, baseURL, apiBasePath, query.Encode())

		data, err := c.httpClient.send(ctx, http.MethodGet, urlStr, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to send: %w", err)
		}

		res := &model.Blogs{}
		err = json.Unmarshal(data, res)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal: %w", err)
		}

		img := regexp.MustCompile(`<img[^>]*src="(.*?)"`)
		html := regexp.MustCompile("<[^>]*>")

		for _, d := range res.Data {
			matches := img.FindAllStringSubmatch(d.Text, -1)
			for _, match := range matches {
				d.Images = append(d.Images, fmt.Sprintf("%s://%s%s", c.httpProtocol, baseURL, match[1]))
			}

			d.Text = html.ReplaceAllString(d.Text, "")
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
