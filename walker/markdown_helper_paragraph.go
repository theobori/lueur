package walker

import (
	"strings"

	"github.com/yuin/goldmark/ast"
)

// Keep references in prose, but avoid repeating standalone menu entries.
func (w *Walker) walkParagraphLines(node ast.Node) (string, error) {
	_, isTopLevel := node.Parent().(*ast.Document)
	if w.options.ReferencePosition() != AfterBlocks || !isTopLevel {
		return w.walkIteratorHelper(node)
	}

	var result, line strings.Builder
	references := 0
	hasText := false
	flush := func() {
		if references != 1 || hasText {
			result.WriteString(line.String())
		}
		line.Reset()
		references = 0
		hasText = false
	}

	for child := node.FirstChild(); child != nil; child = child.NextSibling() {
		s, err := w.Walk(child)
		if err != nil {
			return "", err
		}
		line.WriteString(s)
		switch child.(type) {
		case *ast.Link, *ast.AutoLink, *ast.Image:
			references++
		default:
			if strings.TrimSpace(s) != "" {
				hasText = true
			}
		}
		if text, ok := child.(*ast.Text); ok && (text.SoftLineBreak() || text.HardLineBreak()) {
			flush()
		}
	}
	flush()
	return result.String(), nil
}
