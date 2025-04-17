package pkg

import "github.com/gomarkdown/markdown"

func MarkdownToHTML(input []byte) string {
	html := markdown.ToHTML(input, nil, nil)
	return string(html)
}
