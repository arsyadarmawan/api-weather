package http

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"openweather/common"
	"time"
)

type (
	Client struct {
		client *resty.Client
	}
)

var (
	defaultOptions = []ClientOption{
		WithRetryCount(3),
		WithTimeout(1 * time.Minute),
	}
)

func NewClient(baseUrl string, opts ...ClientOption) *Client {
	if parsedUrl, err := url.Parse(baseUrl); err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		panic(fmt.Sprintf("invalid host URL: %s", baseUrl))
	}

	client := &Client{
		client: resty.New().SetBaseURL(baseUrl),
	}
	client.applyOptions(defaultOptions)
	client.applyOptions(opts)
	return client
}

func (c *Client) applyOptions(options []ClientOption) {
	for _, option := range options {
		option(c)
	}
}

func (c Client) SetContext(ctx context.Context) *resty.Request {
	return c.client.R().SetContext(ctx)
}

func (c Client) SetAPIKey(ctx context.Context, apiKey string) *resty.Request {
	return c.SetContext(ctx).SetQueryParam(common.ApiKey, apiKey)
}
