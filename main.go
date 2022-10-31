package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>Markdown Preview Tool</title>
  </head>
  <body>
`
	footer = `
  </body>
</html>
`
)

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

//run coordinate the execution of the program
//takes the filename of the MD to preview,
//it reads the content of input markdown file into a slice of bytes
//by using the FileRead()
//then pass content to ParseContent() func
//which is responsible for converting MD to HTML
//and returns a potential error
func run(filename string) error {
	// Read all the data from the input file and check for errors
	mdContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(mdContent)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	return saveHTML(outName, htmlData)
}

//parseContent parse Markdown content into HTML
//receives a slice of bytes representing the content of the MD file
//returns another slice of bytes with the converted content as HTML
func parseContent(mdContent []byte) []byte {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML
	body := blackfriday.Run(mdContent)
	sanitizedBody := bluemonday.UGCPolicy().SanitizeBytes(body)

	// Create a buffer of bytes to write to file
	var buffer bytes.Buffer

	// Write html to bytes buffer
	buffer.WriteString(header)
	buffer.Write(sanitizedBody)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

//saveHTML wrapper of WriteFile to save the content to a html extension
//it receives the entire html Content to be saved, and html file name
//specified by the parameter savedTo
//with file permission of creating a file if it doesn't exist
//that's both readable and writable by the owner's only, readable by everyone
//returns potential error
func saveHTML(savedTo string, data []byte) error {
	// Write the bytes to the file
	return ioutil.WriteFile(savedTo, data, 0644)
}
