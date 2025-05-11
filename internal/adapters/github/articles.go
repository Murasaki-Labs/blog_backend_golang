package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ArticleMeta struct {
	Slug         string   `json:"slug"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Date         string   `json:"date"`
	Author       string   `json:"author"`
	Tags         []string `json:"tags"`
	CoverImage   string   `json:"coverImage"`
	ReadingTime  string   `json:"readingTime"`
	CanonicalUrl string   `json:"canonicalUrl"`
	OgImage      string   `json:"ogImage"`
}

type StructuredData struct {
	Context  string   `json:"@context"`
	Type     string   `json:"@type"`
	Headline string   `json:"headline"`
	Image    []string `json:"image"`
	Author   struct {
		Type string `json:"@type"`
		Name string `json:"name"`
	} `json:"author"`
	Publisher struct {
		Type string `json:"@type"`
		Name string `json:"name"`
		Logo struct {
			Type string `json:"@type"`
			Url  string `json:"url"`
		} `json:"logo"`
	} `json:"publisher"`
	DatePublished string `json:"datePublished"`
	DateModified  string `json:"dateModified"`
}

func (c *Client) FetchArticlesJSON() ([]ArticleMeta, error) {
	url := fmt.Sprintf("%s/articles.json", baseRawURL)

	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
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

	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
