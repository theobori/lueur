package walker

import (
	"fmt"
	"strings"

	lhtml "github.com/theobori/lueur/html"
	"golang.org/x/net/html"
)

func (w *Walker) walkHTMLFromString(s string) (string, error) {
	reader := strings.NewReader(s)
	node, err := html.Parse(reader)
	if err != nil {
		return "", err
	}

	return w.walkHTML(node)
}

func (w *Walker) walkHTMLTextNode(node *html.Node) (string, error) {
	return node.Data, nil
}

func (w *Walker) walkHTMLP(node *html.Node) (string, error) {
	s, err := w.walkHTMLIteratorHelper(node)
	if err != nil {
		return "", err
	}

	s = "\n" + strings.Trim(s, "\n") + "\n"

	return s, nil
}

func (w *Walker) walkHTMLB(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH1(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH2(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH3(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH4(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH5(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLH6(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLHTML(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLHead(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLBody(node *html.Node) (string, error) {
	return w.walkHTMLIteratorHelper(node)
}

func (w *Walker) walkHTMLDiv(node *html.Node) (string, error) {
	s, err := w.walkHTMLIteratorHelper(node)
	if err != nil {
		return "", err
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkHTMLCenter(node *html.Node) (string, error) {
	s, err := w.walkHTMLIteratorHelper(node)
	if err != nil {
		return "", err
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkHTMLImg(node *html.Node) (string, error) {
	h := lhtml.MapFromAttributes(node.Attr)

	src, hasSrc := h["src"]
	if !hasSrc {
		return "", fmt.Errorf("the 'src' attribute for the 'img' node is mandatory")
	}

	alt, hasAlt := h["alt"]
	var altValue string
	if !hasAlt {
		altValue = ""
	} else {
		altValue = alt.Val
	}

	return w.walkHTMLReferenceHelper(
		altValue,
		src.Val,
	)
}

func (w *Walker) walkHTMLStyle(_ *html.Node) (string, error) {
	return "", nil // Skip style tags since it will never handle styles
}

func (w *Walker) walkHTMLElementNode(node *html.Node) (string, error) {
	switch node.Data {
	case "html":
		return w.walkHTMLHTML(node)
	case "head":
		return w.walkHTMLHead(node)
	case "body":
		return w.walkHTMLBody(node)
	case "h1":
		return w.walkHTMLH1(node)
	case "h2":
		return w.walkHTMLH2(node)
	case "h3":
		return w.walkHTMLH3(node)
	case "h4":
		return w.walkHTMLH4(node)
	case "h5":
		return w.walkHTMLH5(node)
	case "h6":
		return w.walkHTMLH6(node)
	case "b":
		return w.walkHTMLB(node)
	case "p":
		return w.walkHTMLP(node)
	case "div":
		return w.walkHTMLDiv(node)
	case "center":
		return w.walkHTMLCenter(node)
	case "img":
		return w.walkHTMLImg(node)
	case "style":
		return w.walkHTMLStyle(node)
	default:
		return "", fmt.Errorf("unsupported HTML node type: %s", node.Data)
	}
}

func (w *Walker) walkHTML(node *html.Node) (string, error) {
	switch node.Type {
	case html.ElementNode:
		return w.walkHTMLElementNode(node)
	case html.TextNode:
		return w.walkHTMLTextNode(node)
	default:
		return w.walkHTMLIteratorHelper(node)
	}
}

func (w *Walker) WalkHTML(node *html.Node) (string, error) {
	w.ctx.Depth.Add()
	s, err := w.walkHTML(node)
	if err != nil {
		return "", err
	}
	w.ctx.Depth.Remove()

	return s, nil
}
