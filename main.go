package main

import (
	"flag"
	"fmt"
)

func main() {
	urlflag := flag.String("url", "https://github.com/nemopss", "the url you want to build a sitemap")
	flag.Parse()

	fmt.Println(*urlflag)
}

/*
	TODO:
	1. GET the webpage
	2. Parse all the links on the page
	3. Build proper urls with our links
	4. Filter out any links with a different domain
	5. Find all pages (BFS)
	6. Print out XML
*/
