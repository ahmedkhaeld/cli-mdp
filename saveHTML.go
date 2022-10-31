package main

import "io/ioutil"

//saveHTML save the result into a file to transform MD into HTML
//it receives the entire html Content to be saved, and html file name
//specified by the parameter savedFile
//with file permission of creating a file
//that's both readable and writable by the owner's only, readable by everyone
//returns potential error
func saveHTML(savedFile string, htmlContent []byte) error {
	//write the html content to the savedFile
	return ioutil.WriteFile(savedFile, htmlContent, 0644)
}
