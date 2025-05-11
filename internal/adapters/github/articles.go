package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ArticleMeta struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Author      struct {
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"author"`
	Tags           []string       `json:"tags"`
	Difficulty     string         `json:"difficulty"`
	CoverImage     string         `json:"coverImage"`
	ReadingTime    string         `json:"readingTime"`
	CanonicalURL   string         `json:"canonicalUrl"`
	OgImage        string         `json:"ogImage"`
	StructuredData StructuredData `json:"structuredData"`
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
			URL  string `json:"url"`
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

	for i, article := range articles {
		articles[i].StructuredData = generateStructuredData(
			article.Title, article.Date, article.Author.Name, article.CoverImage,
		)
	}

	return articles, nil
}

func (c *Client) FetchArticleJSON(slug string) (*ArticleMeta, error) {
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

	var meta = ArticleMeta{}
	for i, article := range articles {
		articles[i].StructuredData = generateStructuredData(
			article.Title, article.Date, article.Author.Name, article.CoverImage,
		)

		if article.Slug == slug {
			meta = articles[i]
			break
		}
	}

	return &meta, nil
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

func generateStructuredData(title, date, author, coverImage string) StructuredData {
	return StructuredData{
		Context:       "https://schema.org",
		Type:          "BlogPosting",
		Headline:      title,
		Image:         []string{coverImage},
		DatePublished: date,
		DateModified:  date,
		Author: struct {
			Type string `json:"@type"`
			Name string `json:"name"`
		}{
			Type: "Person",
			Name: author,
		},
		Publisher: struct {
			Type string `json:"@type"`
			Name string `json:"name"`
			Logo struct {
				Type string `json:"@type"`
				URL  string `json:"url"`
			} `json:"logo"`
		}{
			Type: "Organization",
			Name: "Murasaki Labs",
			Logo: struct {
				Type string `json:"@type"`
				URL  string `json:"url"`
			}{
				Type: "ImageObject",
				URL:  "https://avatars.githubusercontent.com/u/187413780",
			},
		},
	}
}
