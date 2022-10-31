# MarkDown Preview
MarkDown Preview tool`mdp`, accept the file name of Markdown
file to be previewed as its argument. This tool will perform

four main steps:
1. Read the content of the input Markdown file
2. Use some Go external libraries to parse Markdown and generate valid HTML block
3. Wrap the results with an HTML header and footers
4. Save the buffer to an HTML file that you can view in a browser