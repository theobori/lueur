
# Gophermap renderer

((build badge)[1])[2]

((builtwithnix badge)[3])[4]

The name of the project is lueur, pronounced \lɥœʁ\, which
is a French word for vivid, momentary expression.

This GitHub repository is a KISS project in the form of a
CLI tool. It allows you to convert Markdown and HTML to
gophermap, .gph and .txt files. This project was originally
designed to convert blog posts written in Markdown format.
The (goldmark)[5] project was used to parse the Markdown.

The project was tested with the Gopher servers
(gopher://bitreich.org:70/scm/geomyidae/)[6] and
(gophernicus)[7] and with the Gopher clients (lagrange)[8]
and (Gophie)[9]. I mainly used
(gopher://baud.baby/0/phlog/fs20181102.txt)[10] as a
reference for the gophermap file format.

This project is still under construction, HTML is not
implemented and only the most basic features will be, such
as basic tags, CSS will not be supported.

## Getting started

To start using the tool, simply run the following command.

lueur -help

Below are the CLI options.

Usage of lueur:
  -domain string
        Gopher domain
  -fancy-header
        Write fancy headers (with hashtags as prefix)
  -file-format string
        Output file format ("gophermap", "gph", "txt") (default
"gophermap")
  -filename string
        Read input from a file
  -port int
        Gopher port (default 70)
  -reference-position string
        Used to control where the references are outputed
("after-block", "after-all") (default "after-block")
  -word-wrap-limit int
        Word wrap limit (default 80)

## The Gopher protocol

To understand what Gopher is, I recommend
(AboutGopher.txt)[11] and the (wikipedia page)[12].

## Contribute

If you want to help the project, you can follow the
guidelines in (CONTRIBUTING.md)[13].
[1] https://github.com/theobori/lueur/actions/workflows/build.yml/badge.svg
[2] https://github.com/theobori/lueur/actions/workflows/build.yml
[3] https://builtwithnix.org/badge.svg
[4] https://builtwithnix.org
[5] https://github.com/yuin/goldmark
[6] gopher://bitreich.org:70/scm/geomyidae/
[7] https://github.com/gophernicus/gophernicus
[8] https://gmi.skyjake.fi/lagrange/
[9] https://gophie.org/
[10] gopher://baud.baby/0/phlog/fs20181102.txt
[11] https://github.com/sgolovine/roll-your-gopher/blob/master/AboutGopher.txt
[12] https://en.wikipedia.org/wiki/Gopher_(protocol)
[13] ./CONTRIBUTING.md

