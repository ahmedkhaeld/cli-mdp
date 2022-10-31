package main

import (
	"bytes"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

//header and footer to help generate html browser review
//as the dependent pkg does not have a header and footer
const (
	header = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html; charset=utf-8">​
		<title>Markdown Preview Tool</title>
	</head>
	<body>
`
	footer = `
	</body>
</html>
`
)

//parseContent parse Markdown content into HTML
//receives a slice of bytes representing the content of the MD file
//returns another slice of bytes with the converted content as HTML
func parseContent(mdContent []byte) []byte {

	//parse the md file content through black friday Run func
	html := blackfriday.Run(mdContent)
	//generate valid and safe html
	body := bluemonday.UGCPolicy().SanitizeBytes(html)

	//combine the body with header and footer const
	//to generate the complete HTML content
	//use a buffer of bytes to join html parts
	var buffer bytes.Buffer
	//write html to the bytes buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}
