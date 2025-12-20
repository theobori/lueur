package markdown

import (
	"fmt"
	"strings"

	"github.com/muesli/reflow/wordwrap"
	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/renderer"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	east "github.com/yuin/goldmark/extension/ast"
	"github.com/yuin/goldmark/text"
)

type Renderer struct {
	node    ast.Node
	source  []byte
	depth   int
	options *renderer.Options

	// Hashmap that manage the references output
	referencesQueue []string
}

func NewRendererWithOptions(source []byte, options *renderer.Options) *Renderer {
	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
	)

	p := markdown.Parser()
	node := p.Parse(text.NewReader(source))

	return &Renderer{
		node:            node,
		source:          source,
		depth:           0,
		options:         options,
		referencesQueue: []string{},
	}
}

func NewRenderer(source []byte, domain string) *Renderer {
	return NewRendererWithOptions(source, renderer.NewDefaultOptions(domain))
}

func (r *Renderer) createLineString(
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
	if r.options.WriteGPHStyle {
		s = line.StringGPHFormat()
	} else {
		s = line.String()
	}

	return s, nil
}

func (r *Renderer) visitEmphasis(node ast.Node) (string, error) {
	return r.visitIteratorHelper(node)
}

func (r *Renderer) visitText(node ast.Node) (string, error) {
	text := node.(*ast.Text)

	s := string(text.Value(r.source))

	if text.SoftLineBreak() || text.HardLineBreak() {
		s += "\n"
	}

	return s, nil
}

func (r *Renderer) visitParagraph(node ast.Node) (string, error) {
	s, err := r.visitIteratorHelper(node)
	if err != nil {
		return "", err
	}

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (r *Renderer) visitThematicBreak(_ ast.Node) (string, error) {
	return "", nil
}

func (r *Renderer) visitHeading(node ast.Node) (string, error) {
	s, err := r.visitIteratorHelper(node)
	if err != nil {
		return "", err
	}

	if r.options.WriteFancyHeaders {
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

func (r *Renderer) visitAutoLink(node ast.Node) (string, error) {
	autoLink := node.(*ast.AutoLink)

	destination := string(autoLink.URL(r.source))
	title := destination

	return r.visitReferenceHelper(node, title, destination)
}

func (r *Renderer) visitBlockQuote(node ast.Node) (string, error) {
	s, err := r.visitIteratorHelper(node)
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

func (r *Renderer) visitList(_ ast.Node) (string, error) {
	// TODO: implement lists
	return "", nil
}

func (r *Renderer) visitLink(node ast.Node) (string, error) {
	link := node.(*ast.Link)
	destination := string(link.Destination)
	title := string(link.Title)

	return r.visitReferenceHelper(node, title, destination)
}

func (r *Renderer) visitImage(node ast.Node) (string, error) {
	image := node.(*ast.Image)
	destination := string(image.Destination)
	title := string(image.Title)

	return r.visitReferenceHelper(node, title, destination)
}

func (r *Renderer) visitCodeBlock(node ast.Node) (string, error) {
	s := string(node.Lines().Value(r.source))
	s = strings.Trim(s, "\n")

	// since every gopher clients dont react the same
	// it could be interesting to paste the code into a new text file
	//
	// then we will add a new references the the hashmap

	if node.HasBlankPreviousLines() {
		s = "\n" + s
	}

	s += "\n"

	return s, nil
}

func (r *Renderer) visitDocument(node ast.Node) (string, error) {
	return r.visitIteratorHelper(node)
}

func (r *Renderer) visitTextBlock(node ast.Node) (string, error) {
	return r.visitIteratorHelper(node)
}

func (r *Renderer) visitHTMLBlock(_ ast.Node) (string, error) {
	return "", nil
}

func (r *Renderer) visitTable(_ ast.Node) (string, error) {
	return "", nil
}

func (r *Renderer) visitCodeSpan(node ast.Node) (string, error) {
	return r.visitIteratorHelper(node)
}

func (r *Renderer) visit(node ast.Node) (string, error) {
	switch node.(type) {
	case *ast.Emphasis:
		return r.visitEmphasis(node)
	case *ast.Text:
		return r.visitText(node)
	case *ast.Paragraph:
		return r.visitParagraph(node)
	case *ast.Heading:
		return r.visitHeading(node)
	case *ast.ThematicBreak:
		return r.visitThematicBreak(node)
	case *ast.AutoLink:
		return r.visitAutoLink(node)
	case *ast.Blockquote:
		return r.visitBlockQuote(node)
	case *ast.List:
		return r.visitList(node)
	case *ast.Link:
		return r.visitLink(node)
	case *ast.Image:
		return r.visitImage(node)
	case *ast.CodeBlock, *ast.FencedCodeBlock:
		return r.visitCodeBlock(node)
	case *ast.Document:
		return r.visitDocument(node)
	case *ast.TextBlock:
		return r.visitTextBlock(node)
	case *ast.HTMLBlock:
		return r.visitHTMLBlock(node)
	case *ast.CodeSpan:
		return r.visitCodeSpan(node)
	case *east.Table:
		return r.visitTable(node)
	default:
		return "", fmt.Errorf("unsupported node type: %s", node.Kind().String())
	}
}

func (r *Renderer) formatDepthOneText(s string) (string, error) {
	s = strings.TrimRight(s, "\n")

	s = wordwrap.String(s, r.options.WordWrapLimit)

	sDest := ""
	linesRaw := strings.SplitSeq(s, "\n")
	for lineRaw := range linesRaw {
		// prevention substitutions
		//
		// replace tabs with four spaces
		lineRaw = strings.ReplaceAll(lineRaw, "\t", "    ")
		// remove antislash at the end
		lineRaw = strings.TrimRight(lineRaw, "\\")

		lineString, err := r.createLineString(
			gophermap.ItemTypeInlineText,
			lineRaw,
			"/",
			r.options.Domain,
			r.options.Port,
		)
		if err != nil {
			return "", err
		}

		sDest += lineString + "\n"
	}

	return sDest, nil
}

func (r *Renderer) Visit(node ast.Node) (string, error) {
	r.depth += 1

	s, err := r.visit(node)
	if err != nil {
		return "", err
	}

	r.depth -= 1

	// the string result at depth 1 should always be gophermap inline text
	// since refs are processed after
	if r.depth == 1 {
		s, err = r.formatDepthOneText(s)
		if err != nil {
			return "", err
		}
	}

	// output the references by using the dedicated Lines
	//
	// depth 0 -> document/root node
	if len(r.referencesQueue) > 0 &&
		(r.options.ReferencesPosition == renderer.AfterBlocks && r.depth == 1) ||
		(r.options.ReferencesPosition == renderer.AfterTraverse && r.depth == 0) {
		for _, line := range r.referencesQueue {
			s += line + "\n"
		}

		r.referencesQueue = nil
	}

	return s, nil
}

func (r *Renderer) VisitFromRoot() (string, error) {
	return r.Visit(r.node)
}

func (r *Renderer) Render() (string, error) {
	return r.VisitFromRoot()
}
