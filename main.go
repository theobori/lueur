package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/walker"
	"github.com/theobori/lueur/walker/markdown"
)

func main() {
	var (
		err                     error
		filename                string
		options                 walker.Options
		referencePositionString string
		fileFormatString        string
	)

	flag.StringVar(
		&filename,
		"filename",
		"",
		"Read input from a file",
	)
	// walker options as CLI flags
	flag.StringVar(
		&options.Domain,
		"domain",
		"",
		"Gopher domain",
	)
	flag.IntVar(
		&options.Port,
		"port",
		gophermap.DefaultGopherPort,
		"Gopher port",
	)
	flag.IntVar(
		&options.WordWrapLimit,
		"word-wrap-limit",
		80,
		"Word wrap limit",
	)
	flag.BoolVar(
		&options.WriteFancyHeader,
		"fancy-header",
		false,
		"Write fancy headers (with hashtags as prefix)",
	)
	flag.StringVar(
		&fileFormatString,
		"file-format",
		"gophermap",
		"Output file format ('gophermap', 'gph')",
	)
	flag.StringVar(
		&referencePositionString,
		"reference-position",
		"after-block",
		"Used to control where the references are outputed ('after-block', 'after-all')",
	)

	flag.Parse()

	if options.Port < 0 {
		log.Fatalln("the port must be a positive integer")
	}
	if options.WordWrapLimit < 50 {
		log.Fatalln("the word wrap limit must be at least 50")
	}
	if options.Domain == "" {
		log.Fatalln("you must set the domain flag")
	}

	referencePosition, err := walker.NewOutputPositionFromString(referencePositionString)
	if err != nil {
		log.Fatalln(err)
	}

	options.ReferencePosition = referencePosition

	fileFormat, err := gophermap.NewFileFormatFromString(fileFormatString)
	if err != nil {
		log.Fatalln(err)
	}

	options.FileFormat = fileFormat

	var source []byte
	if filename == "" {
		source, err = io.ReadAll(os.Stdin)
	} else {
		source, err = os.ReadFile(filename)
	}

	if err != nil {
		log.Fatalln(err)
	}

	w := markdown.NewWalkerWithOptions(source, &options)

	destination, err := w.WalkFromRoot()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(destination)
}
