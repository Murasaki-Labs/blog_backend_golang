package github

import (
	"context"
	"net/http"
)

const (
	baseRawURL = "https://raw.githubusercontent.com/Murasaki-Labs/blog_articles_static/main"
)

type Client struct {
	ctx        context.Context
	httpClient *http.Client
}

func NewClient(ctx context.Context) *Client {
	return &Client{
		ctx:        ctx,
		httpClient: &http.Client{},
	}
}
