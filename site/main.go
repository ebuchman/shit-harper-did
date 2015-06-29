package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ebuchman/shit-harper-did/site/Godeps/_workspace/src/github.com/russross/blackfriday"
)

func main() {
	input, err := ioutil.ReadFile("shit-harper-did.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	buf.WriteString(`<!DOCTYPE html>
<html>
<head>
<link rel="stylesheet" href="style.css">
</head>

<body>
`)
	buf.Write(blackfriday.MarkdownBasic(input))
	buf.WriteString(`
</body>
</html>
`)

	if err := ioutil.WriteFile("index.html", buf.Bytes(), 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
