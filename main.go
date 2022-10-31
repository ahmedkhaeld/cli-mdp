package main

import (
	"flag"
	"os"
)

func main() {
	//parse flags
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	//if user did not provide input file, show usage
	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

}
