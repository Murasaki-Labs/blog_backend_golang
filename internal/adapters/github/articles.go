package github

import (
	"encoding/json"
	"fmt"
	"io"
)

type ArticleMeta struct {
	Slug         string `json:"slug"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	PreviewImage string `json:"previewImage"`
	Date         string `json:"date"`
}

func (c *Client) FetchArticlesJSON() ([]ArticleMeta, error) {
	url := fmt.Sprintf("%s/articles.json", baseRawURL)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var articles []ArticleMeta
	if err := json.NewDecoder(resp.Body).Decode(&articles); err != nil {
		return nil, err
	}
	return articles, nil
}

func (c *Client) FetchMarkdown(slug string) ([]byte, error) {
	fmt.Println("FETCH")
	url := fmt.Sprintf("%s/articles/%s/index.md", baseRawURL, slug)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
