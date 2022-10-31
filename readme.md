# MarkDown Preview
MarkDown Preview tool`mdp`, accept the file name of Markdown
file to be previewed as its argument. This tool will perform

four main steps:
1. Read the content of the input Markdown file
2. Use some Go external libraries to parse Markdown and generate valid HTML block
3. Wrap the results with an HTML header and footers
4. Save the buffer to an HTML file that you can view in a browser

---
* Use the `blackfriday` pkg generate the content based on the input
Markdown, but it doesn't include the HTML header and footer required to view it in a browser.
* Use the `bluemonday` pkg to sanitize the html content to be valid and safe