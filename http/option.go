package http

import "time"

type ClientOption func(*Client)

func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.client.SetTimeout(timeout)
	}
}

func WithRetryCount(count int) ClientOption {
	return func(c *Client) {
		c.client.SetRetryCount(count)
	}
}

func WithRetryWaitTime(waitTime time.Duration) ClientOption {
	return func(c *Client) {
		c.client.SetRetryWaitTime(waitTime)
	}
}

func WithRetryMaxWaitTime(maxWaitTime time.Duration) ClientOption {
	return func(c *Client) {
		c.client.SetRetryMaxWaitTime(maxWaitTime)
	}
}
