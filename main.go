package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/renderer"
	"github.com/theobori/lueur/renderer/markdown"
)

func main() {
	var (
		err      error
		filename string
		options  renderer.Options
	)

	flag.StringVar(&filename, "filename", "", "Read input from a file")

	// Renderer options as CLI flags
	flag.StringVar(&options.Domain, "domain", "", "Gopher domain")
	flag.IntVar(&options.Port, "port", gophermap.DefaultGopherPort, "Gopher port")
	flag.IntVar(&options.WordWrapLimit, "word-wrap-limit", 80, "Word wrap limit")
	flag.BoolVar(&options.WriteFancyHeaders, "fancy-header", false, "Write fancy headers (with hashtags as prefix)")
	flag.BoolVar(&options.WriteGPHStyle, "gph-style", false, "Use gph style instead of gophermap")

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

	var source []byte
	if filename == "" {
		source, err = io.ReadAll(os.Stdin)
	} else {
		source, err = os.ReadFile(filename)
	}

	if err != nil {
		log.Fatalln(err)
	}

	r := markdown.NewRendererWithOptions(source, &options)

	destination, err := r.Render()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(destination)
}
