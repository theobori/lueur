package walker

import (
	"strings"

	"github.com/theobori/lueur/gophermap"
	"github.com/yuin/goldmark/ast"
)

// TODO: add memoization
func (w *Walker) ContainsOnlyRefs(node ast.Node) (bool, error) {
	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		switch c.(type) {
		case *ast.AutoLink, *ast.Link, *ast.Image:
			continue
		case *ast.Text, *ast.TextBlock:
			s, err := w.Walk(c)
			if err != nil {
				return false, err
			}

			s = strings.TrimSpace(s)
			if len(s) > 0 {
				return false, nil
			}
		default:
			return false, nil
		}
	}

	return true, nil
}

// TODO: add memoization
func (w *Walker) ParentsContainsOnlyRefs(node ast.Node) (bool, error) {
	curr := node.Parent()
	for {
		_, isDocument := curr.(*ast.Document)
		if isDocument {
			break
		}

		c, err := w.ContainsOnlyRefs(curr)
		if err != nil {
			return false, err
		}
		if !c {
			return false, nil
		}
		curr = curr.Parent()
	}

	return true, nil
}

func (w *Walker) referenceDescription(node ast.Node, title string, destination string) (string, error) {
	if title != "" {
		return title, nil
	}

	description, err := w.walkIteratorHelper(node)
	if err != nil {
		return "", err
	}

	if description == "" {
		return destination, nil
	}

	return description, nil
}

func (w *Walker) walkReferenceHelper(node ast.Node, title string, destination string) (string, error) {
	description, err := w.referenceDescription(node, title, destination)
	if err != nil {
		return "", err
	}

	line, err := w.referenceLine(description, destination)
	if err != nil {
		return "", err
	}

	inlineAnswer := w.processReferenceLineEdgeCases(line, destination)

	_, isAutoLink := node.(*ast.AutoLink)
	if isAutoLink &&
		w.options.FileFormat() == gophermap.FileFormatTxt &&
		w.options.ReferencePosition() == AfterTraverse {
		return description, nil
	}

	w.ctx.ReferencesQueue = append(w.ctx.ReferencesQueue, *line)

	return inlineAnswer, nil
}
