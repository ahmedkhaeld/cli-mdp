# MarkDown Preview
MarkDown Preview tool`mdp`, that convert .md extension to .html extension that can be viewed in a browser

How? in four main steps:
1. Read the content of the input Markdown file
2. Use some Go external libraries to parse Markdown and generate valid HTML block
3. Wrap the results with an HTML header and footers
4. Save the buffer to an HTML file that you can view in a browser

---
* Use the `blackfriday` pkg generate the content based on the input
Markdown, but it doesn't include the HTML header and footer required to view it in a browser.
* Use the `bluemonday` pkg to sanitize the html content to be valid and safe<br>
`- go get github.com/russross/blackfriday/v2` <br>
`- go get github.com/microcosm-cc/bluemonday`

---
* Auto-Preview Feature

use those commands to build and execute the tool:<br>
the preview file opening automatically in the browser.
``` 
go build -o mdp
./mdp -file README.md
```
