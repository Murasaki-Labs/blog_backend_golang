package github

import (
	"net/http"
)

const (
	baseRawURL = "https://raw.githubusercontent.com/Murasaki-Labs/blog_articles_static/main"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}
