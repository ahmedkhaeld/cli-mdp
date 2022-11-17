package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"time"
)

const (
	defaultTemplate = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html; charset=utf-8">
    <title>{{ .Title }}</title>
  </head>
  <body>
{{ .Body }}
  </body>
</html>
`
)

// content type represents the HTML content tto add into the template
type content struct {
	Title string
	Body  template.HTML
}

func main() {
	// Parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	skipPreview := flag.Bool("s", false, "Skip auto-preview")
	tFname := flag.String("t", "", "Alternate template name")
	flag.Parse()

	// If user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename, *tFname, os.Stdout, *skipPreview); err != nil {
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
func run(filename string, tFname string, out io.Writer, skipPreview bool) error {
	// Read all the data from the input file and check for errors
	mdContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(mdContent, tFname)
	if err != nil {
		return err
	}

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

	defer os.Remove(outName)

	return preview(outName)
}

//parseContent parse Markdown content into HTML
//receives a slice of bytes representing the content of the MD file
//returns another slice of bytes with the converted content as HTML
func parseContent(mdContent []byte, tFname string) ([]byte, error) {
	// Parse the markdown file through blackfriday and bluemonday
	// to generate a valid and safe HTML
	body := blackfriday.Run(mdContent)
	sanitizedBody := bluemonday.UGCPolicy().SanitizeBytes(body)

	// Parse the contents of the defaultTemplate const into a new Template
	t, err := template.New("mdp").Parse(defaultTemplate)
	if err != nil {
		return nil, err
	}

	// If user provided alternate template file, replace template
	if tFname != "" {
		t, err = template.ParseFiles(tFname)
		if err != nil {
			return nil, err
		}
	}

	// Instantiate the content type, adding predefined title and body
	c := content{
		Title: "Markdown Preview Tool",
		Body:  template.HTML(sanitizedBody),
	}
	// Create a buffer of bytes to store the template execution's result
	var buffer bytes.Buffer

	// Execute the template with the content type,
	//inject c into the template and write the results to the buffer
	if err := t.Execute(&buffer, c); err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
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
	err = exec.Command(cPath, cParams...).Run()

	//add delay for preview to be displayed before cleaning up
	time.Sleep(2 * time.Second)

	return err
}
