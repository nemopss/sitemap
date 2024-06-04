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
