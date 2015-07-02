package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ebuchman/shit-harper-did/site/Godeps/_workspace/src/github.com/russross/blackfriday"
)

// simple script to generate index.html from shit-harper-did.md
// run it from the `site` directory with `go run main.go`
// includes the style file and the "fork me on github" ribbon
func main() {
	input, err := ioutil.ReadFile("shit-harper-did.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	buf := new(bytes.Buffer)
	// write header and "fork me on github" ribbon (https://codepo8.github.io/css-fork-on-github-ribbon/)
	buf.WriteString(`<!DOCTYPE html>
<html>
<head>
	<link rel="stylesheet" type="text/css" href="style.css">
	<link href='http://fonts.googleapis.com/css?family=Neuton:200,300,400' rel='stylesheet' type='text/css'>
	<link href='http://fonts.googleapis.com/css?family=Montserrat:400,700' rel='stylesheet' type='text/css'>
</head>

<body>

<style>#forkongithub a{background:#2288ff;color:#fff;text-decoration:none;font-family:arial,sans-serif;text-align:center;font-weight:bold;padding:5px 40px;font-size:1rem;line-height:2rem;position:relative;transition:0.5s;}#forkongithub a:hover{background:#0000ff;color:#fff;}#forkongithub a::before,#forkongithub a::after{content:"";width:100%;display:block;position:absolute;top:1px;left:0;height:1px;background:#fff;}#forkongithub a::after{bottom:1px;top:auto;}@media screen and (min-width:800px){#forkongithub{position:fixed;display:block;top:0;right:0;width:200px;overflow:hidden;height:200px;z-index:9999;}#forkongithub a{width:200px;position:absolute;top:60px;right:-60px;transform:rotate(45deg);-webkit-transform:rotate(45deg);-ms-transform:rotate(45deg);-moz-transform:rotate(45deg);-o-transform:rotate(45deg);box-shadow:;}}</style><span id="forkongithub"><a href="https://github.com/ebuchman/shit-harper-did">Fork me on GitHub</a></span>
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
