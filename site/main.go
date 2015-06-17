package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ebuchman/shit-harper-did/Godeps/_workspace/src/github.com/russross/blackfriday"
)

func main() {
	input, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := blackfriday.MarkdownBasic(input)
	if err := ioutil.WriteFile("index.html", output, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
