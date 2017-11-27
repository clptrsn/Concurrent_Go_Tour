package main

import (
	"fmt"
	"net/http"
	"io"
	"golang.org/x/net/html"
)

func getLinks(reader io.Reader) []string {
	links := []string{}

	tokenizer := html.NewTokenizer(reader)
	token := tokenizer.Next()
	// Loop through the tokens until End of HTML or error
	for token != html.ErrorToken {
		switch token {
		case html.StartTagToken:
			// Get the current token in the stream
			startToken := tokenizer.Token()

			if startToken.Data == "a" {
				// Loop through attributes of the current <a> tag
				for _, attr := range startToken.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}
				}
			}
		}

		token = tokenizer.Next()
	}

	return links
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	urls := getLinks(resp.Body)

	// TODO: Add in the url base if the child url starts with /
	fmt.Printf("found: %s with %d links\n", url, len(urls))

	for _, u := range urls {
		fmt.Printf("Exploring %s\n", u)
		Crawl(u, depth-1)
	}
	return
}

func main() {
	Crawl("http://golang.org/", 1)
}