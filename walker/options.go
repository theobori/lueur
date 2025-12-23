package walker

import (
	"github.com/theobori/lueur/gophermap"
)

type Options struct {
	// Maximum amount of characters per line
	WordWrapLimit int
	// Control after which entity references will be outputed
	ReferencePosition OutputPosition
	// Gopher site domain
	Domain string
	// Gopher port
	Port int
	// Write fancy headers with hashtags as prefix
	WriteFancyHeader bool
	// Write using the GPH format
	FileFormat gophermap.FileFormat
}

func NewDefaultOptions(domain string) *Options {
	return &Options{
		WordWrapLimit:     80,
		ReferencePosition: AfterBlocks,
		Domain:            domain,
		Port:              gophermap.DefaultGopherPort,
		WriteFancyHeader:  false,
		FileFormat:        gophermap.FileFormatGophermap,
	}
}
