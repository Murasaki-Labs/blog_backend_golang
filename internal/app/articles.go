package app

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/orgs/murasaki-labs/blog-backend/internal/adapters/github"
	"github.com/orgs/murasaki-labs/blog-backend/pkg"
)

func (a *App) GetArticlesList() ([]github.ArticleMeta, error) {
	const cacheKey = "articles_list"

	if cached, found := a.cache.Get(cacheKey); found {
		a.log.Debug("Returning cached article list")
		return cached.([]github.ArticleMeta), nil
	}

	a.log.Debug("Fetching article list from GitHub")
	articles, err := a.clients.GitHub().FetchArticlesJSON()
	if err != nil {
		return nil, err
	}

	a.cache.Set(cacheKey, articles, 1*time.Minute)
	return articles, nil
}

func (a *App) GetArticleBySlugHTML(slug string) ([]byte, error) {
	cacheKey := fmt.Sprintf("article:html:%s", slug)

	if cached, found := a.cache.Get(cacheKey); found {
		a.log.Debug("Returning cached article HTML", "slug", slug)
		return cached.([]byte), nil
	}

	a.log.Debug("Fetching article markdown", "slug", slug)
	md, err := a.clients.GitHub().FetchMarkdown(slug)
	if err != nil {
		return nil, err
	}

	html := []byte(pkg.MarkdownToHTML(md))

	a.cache.Set(cacheKey, html, defaultCacheExpiration)

	return html, nil
}

func (a *App) GetArticleBySlug(slug string) ([]byte, error) {
	cacheKey := fmt.Sprintf("article:json:%s", slug)

	if cached, found := a.cache.Get(cacheKey); found {
		a.log.Debug("Returning cached article JSON", "slug", slug)
		return cached.([]byte), nil
	}

	a.log.Debug("Fetching article json", "slug", slug)
	meta, err := a.clients.GitHub().FetchArticleJSON(slug)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(meta)
	if err != nil {
		return nil, err
	}

	a.cache.Set(cacheKey, jsonData, defaultCacheExpiration)
	return jsonData, nil
}
