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
	WeatherConfig struct {
		BaseUrl string
		ApiKey  string
	}
)

var (
	defaultOptions = []ClientOption{
		WithRetryCount(3),
		WithTimeout(1 * time.Minute),
	}
)

func NewClient(request WeatherConfig, opts ...ClientOption) *Client {
	if parsedUrl, err := url.Parse(request.BaseUrl); err != nil || parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		panic(fmt.Sprintf("invalid host URL: %s", request.BaseUrl))
	}

	client := &Client{
		client: resty.New().SetBaseURL(request.BaseUrl).SetQueryParam(
			common.ApiKey, request.ApiKey,
		),
	}
	client.applyOptions(defaultOptions)
	client.applyOptions(opts)
	client.SetAPIKey(request.ApiKey)
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

func (c Client) SetAPIKey(apiKey string) *resty.Request {
	return c.client.R().SetQueryParam(common.ApiKey, apiKey)
}
