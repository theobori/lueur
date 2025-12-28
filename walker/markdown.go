package walker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/muesli/reflow/wordwrap"
	"github.com/theobori/lueur/gophermap"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

type Walker struct {
	node    ast.Node
	source  []byte
	options *Options
	ctx     *Context
}

func NewWalkerWithOptions(source []byte, options *Options) *Walker {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	p := markdown.Parser()
	node := p.Parse(text.NewReader(source))

	return &Walker{
		node:    node,
		source:  source,
		ctx:     NewDefaultContext(),
		options: options,
	}
}

func NewWalker(source []byte, domain string) *Walker {
	defaultOptions, _ := NewDefaultOptions(domain)

	return NewWalkerWithOptions(source, defaultOptions)
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

func (w *Walker) walkList(node ast.Node) (string, error) {
	list := node.(*ast.List)

	marker := string(list.Marker)
	items := []string{}
	i := list.Start

	for c := list.FirstChild(); c != nil; c = c.NextSibling() {
		w.ctx.Indentation.Indent()
		line, err := w.Walk(c)
		if err != nil {
			return "", err
		}

		line = strings.Trim(line, "\n")
		line = marker + " " + line

		if list.IsOrdered() {
			line = strconv.Itoa(i) + line
			i += 1
		}

		w.ctx.Indentation.UnIndent()

		line = w.ctx.Indentation.IndentValue() + strings.TrimLeft(line, " ")

		items = append(items, line)
	}

	s := strings.Join(items, "\n")

	if list.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
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

func (w *Walker) walkListItem(node ast.Node) (string, error) {
	items := []string{}

	for c := node.FirstChild(); c != nil; c = c.NextSibling() {
		line, err := w.Walk(c)
		if err != nil {
			return "", err
		}

		line = strings.Trim(line, "\n")

		items = append(items, line)
	}

	return strings.Join(items, "\n"), nil
}

func (w *Walker) walkDocument(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walkTextBlock(node ast.Node) (string, error) {
	return w.walkIteratorHelper(node)
}

func (w *Walker) walkHTMLBlock(_ ast.Node) (string, error) {
	// TODO: implements
	return "", nil
}

func (w *Walker) walkTable(_ ast.Node) (string, error) {
	// TODO: implements
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
	case *ast.ListItem:
		return w.walkListItem(node)
	case *east.Table:
		return w.walkTable(node)
	default:
		return "", fmt.Errorf("unsupported node type: %s", node.Kind().String())
	}
}

func (w *Walker) formatDepthOneText(s string) (string, error) {
	s = strings.TrimRight(s, "\n")

	if w.options.ReferencePosition() == AfterTraverse && s == "" {
		return "", nil
	}

	s = wordwrap.String(s, w.options.WordWrapLimit())

	sDest := ""
	linesRaw := strings.SplitSeq(s, "\n")

	for lineRaw := range linesRaw {
		// prevention substitutions
		//
		// replace tabs with four spaces
		lineRaw = strings.ReplaceAll(lineRaw, "\t", "    ")
		// remove antislash at the end
		lineRaw = strings.TrimRight(lineRaw, "\\")

		line := gophermap.Line{
			ItemType:    gophermap.ItemTypeInlineText,
			Description: lineRaw,
			Path:        "/",
			Domain:      w.options.Domain(),
			Port:        w.options.Port()}

		sDest += line.StringFromFileFormat(w.options.FileFormat()) + "\n"
	}

	return sDest, nil
}

func (w *Walker) Walk(node ast.Node) (string, error) {
	w.ctx.Depth.Add()

	s, err := w.walk(node)
	if err != nil {
		return "", err
	}

	w.ctx.Depth.Remove()

	// the string result at depth 1 should always be gophermap inline text
	// since refs are processed after
	if w.ctx.Depth.Value() == 1 {
		s, err = w.formatDepthOneText(s)
		if err != nil {
			return "", err
		}
	}

	// output the references by using the dedicated Lines
	//
	// depth 0 -> document/root node
	if len(w.ctx.ReferencesQueue) > 0 &&
		(w.options.ReferencePosition() == AfterBlocks && w.ctx.Depth.Value() == 1) ||
		(w.options.ReferencePosition() == AfterTraverse && w.ctx.Depth.Value() == 0) {
		for _, line := range w.ctx.ReferencesQueue {
			s += line.StringFromFileFormat(w.options.FileFormat()) + "\n"
		}

		w.ctx.ReferencesQueue = nil
	}

	return s, nil
}

func (w *Walker) WalkFromRoot() (string, error) {
	// w.node.Dump(w.source, 0)
	return w.Walk(w.node)
}
