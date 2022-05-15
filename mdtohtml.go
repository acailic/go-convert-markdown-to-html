package main

import (
	"fmt"
	"github.com/gomarkdown/markdown" 
	"io/ioutil"
	"log"
)

func main() {
	file := "test.md"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	html := markdown.ToHTML(content, nil, nil)
	fmt.Println(string(html))
	fileOut := "test.html"
	writeErr := ioutil.WriteFile(fileOut,html,0644)
	if writeErr != nil {
		log.Fatalf("Couldnt write file: %s.",writeErr)
	}
}
