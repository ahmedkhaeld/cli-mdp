package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

//run coordinate the execution of the program
//takes the filename of the MD to preview,
//it reads the content of input markdown file into a slice of bytes
//by using the FileRead()
//then pass content to ParseContent() func
//which is responsible for converting MD to HTML
//and returns a potential error

func run(fn string) error {
	//read the file from the cmd
	mdContent, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}
	// parse the file
	htmlContent := parseContent(mdContent)

	savedFile := fmt.Sprintf("%s.html", filepath.Base(fn))

	//save the parsed content into a html file if error nil
	return saveHTML(savedFile, htmlContent)

}
