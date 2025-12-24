# Gophermap renderer

[![build badge](https://github.com/theobori/lueur/actions/workflows/build.yml/badge.svg)](https://github.com/theobori/lueur/actions/workflows/build.yml)

[![builtwithnix badge](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

The name of the project is lueur, pronounced \lɥœʁ\, which is a French word for vivid, momentary expression.

This GitHub repository is a KISS project in the form of a CLI tool. It allows you to convert Markdown and HTML to gophermap, .gph and .txt files. This project was originally designed to convert blog posts written in Markdown format. The [goldmark](https://github.com/yuin/goldmark) project was used to parse the Markdown.

The project was tested with the Gopher servers [gopher://bitreich.org:70/scm/geomyidae/](gopher://bitreich.org:70/scm/geomyidae/) and [gophernicus](https://github.com/gophernicus/gophernicus) and with the Gopher clients [lagrange](https://gmi.skyjake.fi/lagrange/) and [Gophie](https://gophie.org/). I mainly used [gopher://baud.baby/0/phlog/fs20181102.txt](gopher://baud.baby/0/phlog/fs20181102.txt) as a reference for the gophermap file format.

This project is still under construction, HTML is not implemented and only the most basic features will be, such as basic tags, CSS will not be supported.

## Getting started

To start using the tool, simply run the following command.

```bash
lueur -help
```

Below are the CLI options.

```text
Usage of lueur:
  -domain string
    	Gopher domain
  -fancy-header
    	Write fancy headers (with hashtags as prefix)
  -file-format string
    	Output file format ("gophermap", "gph", "txt") (default "gophermap")
  -filename string
    	Read input from a file
  -port int
    	Gopher port (default 70)
  -reference-position string
    	Used to control where the references are outputed ("after-block", "after-all") (default "after-block")
  -word-wrap-limit int
    	Word wrap limit (default 80)
```

## The Gopher protocol

To understand what Gopher is, I recommend [AboutGopher.txt](https://github.com/sgolovine/roll-your-gopher/blob/master/AboutGopher.txt) and the [wikipedia page](https://en.wikipedia.org/wiki/Gopher_(protocol)).

## Contribute

If you want to help the project, you can follow the guidelines in [CONTRIBUTING.md](./CONTRIBUTING.md).

