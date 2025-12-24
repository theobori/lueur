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
		referencePositionString string
		fileFormatString        string
		wordWrapLimit           int
		domain                  string
		port                    int
		writeFancyHeader        bool
	)

	flag.StringVar(
		&filename,
		"filename",
		"",
		"Read input from a file",
	)
	flag.StringVar(
		&domain,
		"domain",
		"",
		"Gopher domain",
	)
	flag.IntVar(
		&port,
		"port",
		gophermap.DefaultGopherPort,
		"Gopher port",
	)
	flag.IntVar(
		&wordWrapLimit,
		"word-wrap-limit",
		80,
		"Word wrap limit",
	)
	flag.BoolVar(
		&writeFancyHeader,
		"fancy-header",
		false,
		"Write fancy headers (with hashtags as prefix)",
	)
	flag.StringVar(
		&fileFormatString,
		"file-format",
		"gophermap",
		"Output file format (\"gophermap\", \"gph\", \"txt\")",
	)
	flag.StringVar(
		&referencePositionString,
		"reference-position",
		"after-block",
		"Used to control where the references are outputed (\"after-block\", \"after-all\")",
	)

	flag.Parse()

	referencePosition, err := walker.NewOutputPositionFromString(referencePositionString)
	if err != nil {
		log.Fatalln(err)
	}

	fileFormat, err := gophermap.NewFileFormatFromString(fileFormatString)
	if err != nil {
		log.Fatalln(err)
	}

	options, err := walker.NewOptions(
		wordWrapLimit,
		referencePosition,
		domain,
		port,
		writeFancyHeader,
		fileFormat,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var source []byte
	if filename == "" {
		source, err = io.ReadAll(os.Stdin)
	} else {
		source, err = os.ReadFile(filename)
	}

	if err != nil {
		log.Fatalln(err)
	}

	w := markdown.NewWalkerWithOptions(source, options)

	destination, err := w.WalkFromRoot()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(destination)
}
