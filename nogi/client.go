package nogi

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/kqns91/nogi-sdk-go/nogi/model"
	"golang.org/x/exp/slog"
)

type Client interface {
	GetAllMembers(ctx context.Context) ([]*model.Member, error)
	GetBlogs(ctx context.Context, offset, limit int) ([]*model.Blog, error)
	GetAllBlogs(ctx context.Context) ([]*model.Blog, error)
	GetComments(ctx context.Context, blogID string, offset, limit int) ([]*model.Comment, error)
	GetAllComments(ctx context.Context, blogID string) ([]*model.Comment, error)
	GetHandShakes(ctx context.Context, offset, limit int) ([]*model.HandShake, error)
	GetAllHandShakes(ctx context.Context) ([]*model.HandShake, error)
	GetLives(ctx context.Context, offset, limit int) ([]*model.Live, error)
	GetAllLives(ctx context.Context) ([]*model.Live, error)
	GetMovies(ctx context.Context, offset, limit int) ([]*model.Movie, error)
	GetAllMovies(ctx context.Context) ([]*model.Movie, error)
	GetNews(ctx context.Context, date time.Time) ([]*model.News, error)
	GetSchedules(ctx context.Context, date time.Time) ([]*model.Schedule, error)
}

const (
	baseURL       = "www.nogizaka46.com"
	apiBasePath   = baseURL + "s/n46/api/list"
	protocolHTTP  = "http"
	protocolHTTPS = "https"
)

type client struct {
	logger       Logger
	httpProtocol string
	httpClient   *httpClient
}

type option func(*client)

func New(opts ...option) (Client, error) {
	client := &client{
		logger:       newNilLogger(),
		httpProtocol: protocolHTTPS,
		httpClient: &httpClient{
			base: &http.Client{},
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func WithLogger(logger *slog.Logger) func(*client) {
	return func(c *client) {
		c.logger = logger
	}
}

func WithInsecure() func(*client) {
	return func(c *client) {
		c.httpProtocol = "http"
	}
}

type httpClient struct {
	base *http.Client
}

func (c *httpClient) send(ctx context.Context, method, url string, header map[string]string, body []byte) ([]byte, error) {
	if method == "" {
		return nil, errors.New("method is empty")
	}

	if url == "" {
		return nil, errors.New("url is empty")
	}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	resp, err := c.base.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read all: %w", err)
	}

	res := raw[4 : len(raw)-2]

	return res, nil
}
