package cloudprovider

import "dig/internal/config"

type Client struct {
	name string
}

func New(cfg *config.Config) *Client {
	return &Client{
		name: cfg.ClientName,
	}
}

func (c Client) Name() string {
	return c.name
}
