package walker

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/internal/common"
	"github.com/yuin/goldmark/ast"
)

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

func (w *Walker) buildInlineAnswer(line *gophermap.Line, destination string) string {
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

// It handles every edge case depending of the walker options
func (w *Walker) processReferenceEdgeCases(node ast.Node, line *gophermap.Line, destination string) (string, error) {
	baseInlineAnswser := line.Description

	inlineAnswer := w.buildInlineAnswer(line, destination)

	_, isAutoLink := node.(*ast.AutoLink)
	if isAutoLink &&
		w.options.FileFormat() == gophermap.FileFormatTxt &&
		w.options.ReferencePosition() == AfterTraverse {
		return baseInlineAnswser, nil
	}

	w.ctx.ReferencesQueue = append(w.ctx.ReferencesQueue, *line)

	return inlineAnswer, nil
}

func (w *Walker) walkReferenceHelper(node ast.Node, title string, destination string) (string, error) {
	var err error

	line := gophermap.Line{}

	if title != "" {
		line.Description = title
	} else {
		line.Description, err = w.walkIteratorHelper(node)
		if err != nil {
			return "", err
		}
	}

	if line.Description == "" {
		line.Description = destination
	}

	if common.IsURL(destination) {
		u, err := url.Parse(destination)
		if err != nil {
			return "", err
		}

		line.ItemType = gophermap.NewItemTypeFromURL(u)
		line.Domain = u.Host

		line.Port, err = gophermap.PortFromURL(u)
		if err != nil {
			return "", err
		}

		line.Path = gophermap.PathFromURL(u)
	} else {
		line.ItemType = gophermap.NewItemTypeFromPath(destination)
		line.Domain = w.options.Domain()
		line.Port = w.options.Port()
		line.Path = "/" + strings.TrimLeft(destination, "/")
	}

	return w.processReferenceEdgeCases(node, &line, destination)
}
