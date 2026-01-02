package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/theobori/lueur/gophermap"
	"github.com/theobori/lueur/walker"
)

const (
	DirectoryPrefix     = "nix-converter"
	DirectoryOutputName = DirectoryPrefix + "-" + "output"
)

func processFromSource(source []byte, options *walker.Options) (string, error) {
	w := walker.NewWalkerWithOptions(source, options)

	output, err := w.WalkFromRoot()
	if err != nil {
		return "", err
	}

	return output, nil
}

func processFromFilePath(filePath string, options *walker.Options) (string, error) {
	source, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return processFromSource(source, options)
}

// This function should work like a transaction. It means, it will create a temporary directory
// and try to process every Markdown files found. Everything has to succeed,
// otherwise the temporary directory will be removed.
//
// TODO: Improve end user error informations
func processFromDirectoryPath(directoryPath string, outputDirectoryPath string, options *walker.Options) error {
	// Creating the temporary directory
	tDir, err := os.MkdirTemp("", DirectoryPrefix)
	if err != nil {
		return err
	}

	defer os.RemoveAll(tDir)

	err = filepath.WalkDir(directoryPath, func(path string, d fs.DirEntry, err error) error {
		// Skip directories
		if d.IsDir() {
			return nil
		}
		// Skip non Markdown file extension
		ext := filepath.Ext(path)
		if ext != ".md" && ext != ".markdown" {
			return nil
		}

		destinationDirPath := filepath.Join(tDir, filepath.Dir(path))
		destinationFilename := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) + "." + options.FileFormat().String()
		// Create the temporary destination directory
		err = os.MkdirAll(destinationDirPath, os.ModePerm)
		if err != nil {
			return err
		}

		// Process the file path
		output, err := processFromFilePath(path, options)
		if err != nil {
			return fmt.Errorf("error: %s with the file: %s", err, path)
		}

		destinationFilePath := filepath.Join(destinationDirPath, destinationFilename)
		destinationFile, err := os.Create(destinationFilePath)
		if err != nil {
			return err
		}

		_, err = destinationFile.WriteString(output)
		if err != nil {
			return err
		}

		destinationFile.Close()

		log.Printf("The file %s has been written\n", destinationFilePath)

		return nil
	})
	if err != nil {
		return err
	}

	directoryName := filepath.Join(tDir, directoryPath)
	err = os.Rename(directoryName, outputDirectoryPath)
	if err != nil {
		return err
	}

	log.Printf("The directory %s has been created and contains your files", outputDirectoryPath)

	return nil
}

func processFromStdin(options *walker.Options) (string, error) {
	source, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	return processFromSource(source, options)
}

func main() {
	var (
		err                     error
		filePath                string
		directoryPath           string
		outputDirectoryPath     string
		referencePositionString string
		fileFormatString        string
		wordWrapLimit           int
		domain                  string
		port                    int
		writeFancyHeader        bool
		pathPrefix              string
	)

	flag.StringVar(
		&filePath,
		"file",
		"",
		"Read input from a file",
	)
	flag.StringVar(
		&pathPrefix,
		"path-prefix",
		"",
		"Prefix applied to the reference paths",
	)
	flag.StringVar(
		&directoryPath,
		"directory",
		"",
		"Read inputs from a directory, it will recursively take each markdown files as separated inputs",
	)
	flag.StringVar(
		&outputDirectoryPath,
		"output-directory",
		DirectoryOutputName,
		"It specifies an output directory name when -directory-path is used",
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
		pathPrefix,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var output string
	if filePath != "" {
		output, err = processFromFilePath(filePath, options)
	} else if directoryPath != "" {
		err = processFromDirectoryPath(directoryPath, outputDirectoryPath, options)
	} else {
		output, err = processFromStdin(options)
	}

	if err != nil {
		log.Fatalln(err)
	}

	if output != "" {
		fmt.Println(output)
	}
}
