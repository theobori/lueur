
# Gopherspace renderer

((build badge)[1])[2]

((builtwithnix badge)[3])[4]

The name of the project is lueur, pronounced \lɥœʁ\, which
is a French word for vivid, momentary expression.

This GitHub repository is a KISS project in the form of a
CLI tool. It allows you to convert Markdown and HTML to
gophermap, .gph and .txt files. It was originally designed
to convert blog posts written in Markdown format.

## Getting started

To start using the tool, simply run the following command.

lueur -help

There are no CLI shortcuts like -f or -d by choice, as I
prefer to keep them as explicit as possible.

## How it works

The way the project works is deliberately very simple: I
retrieve the text in Markdown format, which can contain
HTML. The text is then passed to the Markdown parser, which
returns an AST that is traversed to produce the final
output. The (goldmark)[5] project was used to parse the
Markdown and the (Go Networking)[6] project for the HTML.
See the (CommonMark specification)[7] to know what is
considered as a HTML block.

Below is a graph summarizing the various stages.

    ┌────────┐
    │1. Input│       1. Usually go by reading a file
    └────┬───┘
         ▼
┌──────────────────┐
│2. Markdown parser│ 2. Parses the given input
└────────┬─────────┘
         ▼
  ┌──────────────┐
  │3. AST walking│   3. Traverses the tree to generate an
  └──────┬───────┘       output
         ▼
    ┌─────────┐
    │4. Output│      4. Prints the output to the process
    └─────────┘          standard output

## Compatibility

It has been tested with the Gopher servers
(gopher://bitreich.org:70/scm/geomyidae/)[8] and
(gophernicus)[9] and with the Gopher clients (lagrange)[10]
and (Gophie)[11]. I mainly used
(gopher://baud.baby/0/phlog/fs20181102.txt)[12] as a
reference for the gophermap file format.

## State

This project is still under construction, only the most
basic HTML features will be implemented, such as basic tags,
CSS will not be supported. It has been designed to support
HTML inside Markdown.

## The Gopher protocol

To understand what Gopher is, I recommend
(AboutGopher.txt)[13] and the (wikipedia page)[14].

## Contribute

If you want to help the project, you can follow the
guidelines in (CONTRIBUTING.md)[15].

## Potential improvements

A potential improvement could be to add a line with
References before writing the references when using the -
reference-position option with the value after-all.

It would also be necessary to detect lines where there are
only references by moving up the tree when a reference is
found, probably with memoization.
[1] https://github.com/theobori/lueur/actions/workflows/build.yml/badge.svg
[2] https://github.com/theobori/lueur/actions/workflows/build.yml
[3] https://builtwithnix.org/badge.svg
[4] https://builtwithnix.org
[5] https://github.com/yuin/goldmark
[6] https://cs.opensource.google/go/x/net/+/master:html/
[7] https://spec.commonmark.org/0.30/#html-blocks
[8] gopher://bitreich.org:70/scm/geomyidae/
[9] https://github.com/gophernicus/gophernicus
[10] https://gmi.skyjake.fi/lagrange/
[11] https://gophie.org/
[12] gopher://baud.baby/0/phlog/fs20181102.txt
[13] https://github.com/sgolovine/roll-your-gopher/blob/master/AboutGopher.txt
[14] https://en.wikipedia.org/wiki/Gopher_(protocol)
[15] ./CONTRIBUTING.md

