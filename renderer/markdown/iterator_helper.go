package markdown

import (
	"github.com/yuin/goldmark/ast"
)

func (r *Renderer) visitIteratorHelper(node ast.Node) (string, error) {
	s := ""

	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		res, err := r.Visit(c)
		if err != nil {
			return "", err
		}

		s += res
	}

	return s, nil
}
