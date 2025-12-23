# Gophermap renderer

[![build](https://github.com/theobori/lueur/actions/workflows/build.yml/badge.svg)](https://github.com/theobori/lueur/actions/workflows/build.yml)

[![built with nix](https://builtwithnix.org/badge.svg)](https://builtwithnix.org)

The name of the project is lueur, pronounced \lɥœʁ\, which is a French word for vivid, momentary expression.

This GitHub repository is a KISS project in the form of a CLI tool. It allows you to convert Markdown and HTML to gophermap and .gph files. This project was originally designed to convert blog posts written in Markdown format. The [goldmark](https://github.com/yuin/goldmark) project was used to parse the Markdown.

The project was tested with the Gopher servers [geomyidae](gopher://bitreich.org:70/scm/geomyidae/) and [gophernicus](https://github.com/gophernicus/gophernicus) and with the Gopher clients [lagrange](https://gmi.skyjake.fi/lagrange/) and [Gophie](https://gophie.org/). I mainly used [fs20181102.txt](gopher://baud.baby/0/phlog/fs20181102.txt) as a reference for the gophermap file format.

## Getting started

To start using the tool, simply run the following command.

```bash
lueur -help
```

## The Gopher protocol

To understand what Gopher is, I recommend [AboutGopher.txt](https://github.com/sgolovine/roll-your-gopher/blob/master/AboutGopher.txt) and the resource [wikipedia page](https://en.wikipedia.org/wiki/Gopher_(protocol)).

## Contribute

If you want to help the project, you can follow the guidelines in [CONTRIBUTING.md](./CONTRIBUTING.md).
