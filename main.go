package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
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
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, os.Stdout, *skipPreview); err != nil {
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
func run(filename string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	mdContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData := parseContent(mdContent)

	//preview the md content through a temporary file
	//use TempFile func that generate a random file name
	//store the file in a system temp directory
	//instead of creating file locally
	tf, err := ioutil.TempFile("", "mdp*.html")
	if err != nil {
		return err
	}
	if err := tf.Close(); err != nil {
		return err
	}
	outName := tf.Name()

	//print the outName to the out
	fmt.Fprintln(out, outName)

	if err := saveHTML(outName, htmlData); err != nil {
		return err
	}

	if skipPreview {
		return nil
	}
	return preview(outName)
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

//preview takes the temp file name as input
//and return error in case it can't open the file
func preview(filename string) error {
	cName := ""
	var cParams []string

	//Define executable based on os
	switch runtime.GOOS {
	case "linux":
		cName = "xdg-open"
	case "windows":
		cName = "cmd.exe"
		cParams = []string{"/C", "start"}
	case "darwin":
		cName = "open"
	default:
		return fmt.Errorf("OS not supported")
	}
	//append filename to parameter list
	cParams = append(cParams, filename)

	//locate executable in $PATH and executes it
	cPath, err := exec.LookPath(cName)
	if err != nil {
		return err
	}

	//open the file using default program
	return exec.Command(cPath, cParams...).Run()
}
