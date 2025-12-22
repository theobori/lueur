package markdown

import (
	"fmt"
	"strings"

	"github.com/muesli/reflow/wordwrap"
	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/walker"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

type Walker struct {
	node    ast.Node
	source  []byte
	depth   int
	options *walker.Options

	// Queue that manage the references output
	referencesQueue []string
}

func NewWalkerWithOptions(source []byte, options *walker.Options) *Walker {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	p := markdown.Parser()
	node := p.Parse(text.NewReader(source))

	return &Walker{
		node:            node,
		source:          source,
		depth:           0,
		options:         options,
		referencesQueue: []string{},
	}
}

func NewWalker(source []byte, domain string) *Walker {
	return NewWalkerWithOptions(source, walker.NewDefaultOptions(domain))
}

func (w *Walker) createLineString(
	itemType gophermap.ItemType, description string,
	destination string, domain string, port int,
) (string, error) {
	line, err := gophermap.NewLine(
		itemType,
		description,
		destination,
		domain,
		port,
	)

	if err != nil {
		return "", nil
	}

	var s string
	if w.options.WriteGPHFormat {
		s = line.StringGPHFormat()
	} else {
		s = line.String()
	}

	return s, nil
}

func (w *Walker) walkEmphasis(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walkText(node ast.Node) (string, error) {
	text := node.(*ast.Text)

	s := string(text.Value(w.source))

	if text.SoftLineBreak() || text.HardLineBreak() {
		s += "\n"
	}

	return s, nil
}

func (w *Walker) walkParagraph(node ast.Node) (string, error) {
	s, err := w.walkIteratorHelper(node)
	if err != nil {
		return "", err
	}

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkThematicBreak(_ ast.Node) (string, error) {
	return "", nil
}

func (w *Walker) walkHeading(node ast.Node) (string, error) {
	s, err := w.walkIteratorHelper(node)
	if err != nil {
		return "", err
	}

	if w.options.WriteFancyHeader {
		heading := node.(*ast.Heading)
		fancyPrefix := strings.Repeat("#", heading.Level)

		sLines := strings.Split(s, "\n")
		for i, sLine := range sLines {
			sLines[i] = fancyPrefix + " " + sLine
		}

		s = strings.Join(sLines, "\n")
	}

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkAutoLink(node ast.Node) (string, error) {
	autoLink := node.(*ast.AutoLink)

	destination := string(autoLink.URL(w.source))
	title := destination

	return w.walkReferenceHelper(node, title, destination)
}

func (w *Walker) walkBlockQuote(node ast.Node) (string, error) {
	s, err := w.walkIteratorHelper(node)
	if err != nil {
		return "", err
	}

	s = "“" + strings.Trim(s, "\n") + "”"

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkList(_ ast.Node) (string, error) {
	// TODO: implement lists
	return "", nil
}

func (w *Walker) walkLink(node ast.Node) (string, error) {
	link := node.(*ast.Link)
	destination := string(link.Destination)
	title := string(link.Title)

	return w.walkReferenceHelper(node, title, destination)
}

func (w *Walker) walkImage(node ast.Node) (string, error) {
	image := node.(*ast.Image)
	destination := string(image.Destination)
	title := string(image.Title)

	return w.walkReferenceHelper(node, title, destination)
}

func (w *Walker) walkCodeBlock(node ast.Node) (string, error) {
	s := string(node.Lines().Value(w.source))
	s = strings.Trim(s, "\n")

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (w *Walker) walkDocument(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walkTextBlock(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walkHTMLBlock(_ ast.Node) (string, error) {
	return "", nil
}

func (w *Walker) walkTable(_ ast.Node) (string, error) {
	return "", nil
}

func (w *Walker) walkCodeSpan(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walk(node ast.Node) (string, error) {
	switch node.(type) {
	case *ast.Emphasis:
		return w.walkEmphasis(node)
	case *ast.Text:
		return w.walkText(node)
	case *ast.Paragraph:
		return w.walkParagraph(node)
	case *ast.Heading:
		return w.walkHeading(node)
	case *ast.ThematicBreak:
		return w.walkThematicBreak(node)
	case *ast.AutoLink:
		return w.walkAutoLink(node)
	case *ast.Blockquote:
		return w.walkBlockQuote(node)
	case *ast.List:
		return w.walkList(node)
	case *ast.Link:
		return w.walkLink(node)
	case *ast.Image:
		return w.walkImage(node)
	case *ast.CodeBlock, *ast.FencedCodeBlock:
		return w.walkCodeBlock(node)
	case *ast.Document:
		return w.walkDocument(node)
	case *ast.TextBlock:
		return w.walkTextBlock(node)
	case *ast.HTMLBlock:
		return w.walkHTMLBlock(node)
	case *ast.CodeSpan:
		return w.walkCodeSpan(node)
	case *east.Table:
		return w.walkTable(node)
	default:
		return "", fmt.Errorf("unsupported node type: %s", node.Kind().String())
	}
}

func (w *Walker) formatDepthOneText(s string) (string, error) {
	s = strings.TrimRight(s, "\n")

	if w.options.ReferencePosition == walker.AfterTraverse && s == "" {
		return "", nil
	}

	s = wordwrap.String(s, w.options.WordWrapLimit)

	sDest := ""
	linesRaw := strings.SplitSeq(s, "\n")
	for lineRaw := range linesRaw {
		// prevention substitutions
		//
		// replace tabs with four spaces
		lineRaw = strings.ReplaceAll(lineRaw, "\t", "    ")
		// remove antislash at the end
		lineRaw = strings.TrimRight(lineRaw, "\\")

		lineString, err := w.createLineString(
			gophermap.ItemTypeInlineText,
			lineRaw,
			"/",
			w.options.Domain,
			w.options.Port,
		)
		if err != nil {
			return "", err
		}

		sDest += lineString + "\n"
	}

	return sDest, nil
}

func (w *Walker) Walk(node ast.Node) (string, error) {
	w.depth += 1

	s, err := w.walk(node)
	if err != nil {
		return "", err
	}

	w.depth -= 1

	// the string result at depth 1 should always be gophermap inline text
	// since refs are processed after
	if w.depth == 1 {
		s, err = w.formatDepthOneText(s)
		if err != nil {
			return "", err
		}
	}

	// output the references by using the dedicated Lines
	//
	// depth 0 -> document/root node
	if len(w.referencesQueue) > 0 &&
		(w.options.ReferencePosition == walker.AfterBlocks && w.depth == 1) ||
		(w.options.ReferencePosition == walker.AfterTraverse && w.depth == 0) {
		for _, line := range w.referencesQueue {
			s += line + "\n"
		}

		w.referencesQueue = nil
	}

	return s, nil
}

func (w *Walker) WalkFromRoot() (string, error) {
	return w.Walk(w.node)
}
