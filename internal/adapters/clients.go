package adapters

import (
	"context"
	"github.com/orgs/murasaki-labs/blog-backend/internal/adapters/github"
)

type Clients struct {
	github *github.Client
}

func MustClients(ctx context.Context) *Clients {
	ghc := github.NewClient(ctx)

	return &Clients{
		github: ghc,
	}
}

func (c *Clients) GitHub() *github.Client {
	return c.github
}
