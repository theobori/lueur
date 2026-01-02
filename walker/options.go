package walker

import (
	"fmt"

	"github.com/theobori/lueur/gophermap"
)

const WordWrapLimitMinimum = 45

type Options struct {
	// Maximum amount of characters per line
	wordWrapLimit int
	// Control after which entity references will be outputed
	referencePosition OutputPosition
	// Gopher site domain
	domain string
	// Gopher port
	port int
	// Write fancy headers with hashtags as prefix
	WriteFancyHeader bool
	// Write using the GPH format
	fileFormat gophermap.FileFormat
	// Prefix for references path
	PathPrefix string
}

func NewOptions(
	wordWrapLimit int,
	referencePosition OutputPosition,
	domain string,
	port int,
	writeFancyHeader bool,
	fileFormat gophermap.FileFormat,
	pathPrefix string,
) (*Options, error) {
	var err error

	o := Options{}

	err = o.SetWordWrapLimit(wordWrapLimit)
	if err != nil {
		return nil, err
	}
	err = o.SetReferencePositionAndFileFormat(referencePosition, fileFormat)
	if err != nil {
		return nil, err
	}

	err = o.SetPort(port)
	if err != nil {
		return nil, err
	}

	err = o.SetDomain(domain)
	if err != nil {
		return nil, err
	}

	o.WriteFancyHeader = writeFancyHeader
	o.PathPrefix = pathPrefix

	return &o, nil
}

func NewDefaultOptions(domain string) (*Options, error) {
	o, err := NewOptions(
		80,
		AfterBlocks,
		domain,
		gophermap.DefaultGopherPort,
		false,
		gophermap.FileFormatGophermap,
		"",
	)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Options) WordWrapLimit() int {
	return o.wordWrapLimit
}

func (o *Options) SetWordWrapLimit(wordWrapLimit int) error {
	if wordWrapLimit < WordWrapLimitMinimum {
		return fmt.Errorf(
			"the word wrap limit must be at least %d",
			WordWrapLimitMinimum,
		)
	}

	o.wordWrapLimit = wordWrapLimit

	return nil
}

func (o *Options) ReferencePosition() OutputPosition {
	return o.referencePosition
}

func (o *Options) SetReferencePositionAndFileFormat(
	referencePosition OutputPosition,
	fileFormat gophermap.FileFormat,
) error {
	if fileFormat == gophermap.FileFormatTxt && referencePosition != AfterTraverse {
		return fmt.Errorf(
			"reference position %s cannot be used with the file format %s",
			referencePosition.String(),
			fileFormat.String(),
		)
	}

	o.referencePosition = referencePosition
	o.fileFormat = fileFormat

	return nil
}

func (o *Options) Domain() string {
	return o.domain
}

func (o *Options) SetDomain(domain string) error {
	if o.fileFormat != gophermap.FileFormatTxt && domain == "" {
		return fmt.Errorf(
			"the domain cannot be empty for the file format %s",
			o.fileFormat.String(),
		)
	}

	o.domain = domain

	return nil
}

func (o *Options) Port() int {
	return o.port
}

func (o *Options) SetPort(port int) error {
	if port < 0 {
		return fmt.Errorf("the port must be a positive integer")
	}

	o.port = port

	return nil
}

func (o *Options) FileFormat() gophermap.FileFormat {
	return o.fileFormat
}
