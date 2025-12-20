package renderer

import "github.com/theobori/lueur/gophermap"

type OutputPosition int

const (
	// The node will be outputed after blocks at depth 0
	AfterBlocks OutputPosition = iota
	// The node will be outputed after the AST has been evaluated
	AfterTraverse
)

type Options struct {
	// Maximum amount of characters per line
	WordWrapLimit int
	// Control after which entity references will be outputed
	ReferencesPosition OutputPosition
	// Gopher site domain
	Domain string
	// Gopher port
	Port int
	// Write fancy headers like '## Header 2'
	WriteFancyHeaders bool
	// Write with bracket style
	WriteGPHStyle bool
}

func NewDefaultOptions(domain string) *Options {
	return &Options{
		WordWrapLimit:      80,
		ReferencesPosition: AfterBlocks,
		Domain:             domain,
		Port:               gophermap.DefaultGopherPort,
		WriteFancyHeaders:  false,
		WriteGPHStyle:      false,
	}
}
