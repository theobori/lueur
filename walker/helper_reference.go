package walker

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
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

// Dirty shit, don't try to understand this function
//
// It's just handling the different edge cases with
// specific walker reference position
func (w *Walker) processReferenceLineEdgeCases(line *gophermap.Line, destination string) string {
	var inlineAnswer string

	switch w.options.ReferencePosition() {
	case AfterTraverse:
		number := len(w.ctx.ReferencesQueue) + 1
		inlineAnswer = fmt.Sprintf("(%s)[%d]", line.Description, number)

		if w.options.FileFormat() == gophermap.FileFormatTxt {
			line.Description = destination
		}

		line.Description = fmt.Sprintf("[%d] %s", number, line.Description)
	case AfterBlocks:
		inlineAnswer = line.Description
	}

	return inlineAnswer
}

func (w *Walker) referenceLine(description string, destination string) (*gophermap.Line, error) {
	line := gophermap.Line{
		Description: description,
	}

	if common.IsURL(destination) {
		u, err := url.Parse(destination)
		if err != nil {
			return nil, err
		}

		line.Port, err = gophermap.PortFromURL(u)
		if err != nil {
			return nil, err
		}

		line.ItemType = gophermap.NewItemTypeFromURL(u)
		line.Domain = u.Host
		line.Path = gophermap.PathFromURL(u)
	} else {
		line.Port = w.options.Port()
		line.ItemType = gophermap.NewItemTypeFromPath(destination)
		line.Domain = w.options.Domain()
		line.Path = "/" + strings.TrimLeft(destination, "/")
	}

	return &line, nil
}
