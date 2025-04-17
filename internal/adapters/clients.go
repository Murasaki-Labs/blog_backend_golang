package adapters

import "github.com/orgs/murasaki-labs/blog-backend/internal/adapters/github"

type Clients struct {
	github *github.Client
}

func MustClients() *Clients {
	ghc := github.NewClient()

	return &Clients{
		github: ghc,
	}
}

func (c *Clients) GitHub() *github.Client {
	return c.github
}
