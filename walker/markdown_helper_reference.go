package walker

import (
	"github.com/theobori/lueur/gophermap"
	"github.com/yuin/goldmark/ast"
)

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
