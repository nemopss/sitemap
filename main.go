package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/nemopss/link"
)

func main() {
	urlflag := flag.String("url", "https://gophercises.com/", "the url you want to build a sitemap")
	flag.Parse()

	fmt.Println(*urlflag)

	resp, err := http.Get(*urlflag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	var hrefs []string

	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)

		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
