package pkg

import (
	"bufio"
	"bytes"

	"github.com/gomarkdown/markdown"
)

func MarkdownToHTML(input []byte) string {
	clean := removeYAMLFrontMatter(input)
	return string(markdown.ToHTML(clean, nil, nil))
}

func removeYAMLFrontMatter(input []byte) []byte {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var (
		skipping bool
		started  bool
		buf      bytes.Buffer
	)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "---" {
			if !started {
				started = true
				skipping = true
				continue
			} else if skipping {
				skipping = false
				continue
			}
		}

		if !skipping {
			buf.WriteString(line + "\n")
		}
	}

	return buf.Bytes()
}
