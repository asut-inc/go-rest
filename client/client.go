package client

import (
	"errors"
	"net/http"
)

// NewClient constructs a new Client which can make requests to the Google Maps
// WebService APIs.
func NewClient(options ...Option) (*Client, error) {
	c := &Client{}
	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// WithHTTPClient configures a Maps API client with a http.Client to make requests
// over.
func WithHTTPClient(c *http.Client) Option {
	return func(client *Client) error {
		client.httpClient = c
		if c == nil {
			return errors.New("client: provider http client is nil")
		}
		return nil
	}
}

// Client may be used to make requests to the Google Maps WebService APIs
type Client struct {
	httpClient *http.Client
}

// Option is the type of constructor options for NewClient(...).
type Option func(*Client) error
