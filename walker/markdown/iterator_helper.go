package markdown

import (
	"github.com/yuin/goldmark/ast"
)

func (w *Walker) walkIteratorHelper(node ast.Node) (string, error) {
	s := ""

	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		res, err := w.Walk(c)
		if err != nil {
			return "", err
		}

		s += res
	}

	return s, nil
}
