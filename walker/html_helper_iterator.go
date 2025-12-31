package walker

import (
	"strings"

	"golang.org/x/net/html"
)

func (w *Walker) walkHTMLIteratorHelper(node *html.Node) (string, error) {
	builder := strings.Builder{}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		s, err := w.WalkHTML(c)
		if err != nil {
			return "", err
		}

		_, err = builder.WriteString(s)
		if err != nil {
			return "", err
		}
	}

	return builder.String(), nil
}
