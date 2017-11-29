package main

import (
	"fmt"
	"net/http"
	"io"
	"golang.org/x/net/html"
	"regexp"
	"strings"
	"time"
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

var urlMaps map[string][]string


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	if depth <= 0 {
		return
	}	

	urlRegex, _ := regexp.Compile(`(?P<Scheme>http(?:s|):\/\/)(?P<Path>.*)(?:\/|$)`)

	isValidUrl := urlRegex.MatchString(url)
	if isValidUrl == false {
		return
	}

	match := urlRegex.FindStringSubmatch(url)
	baseUrl := match[1] + match[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	urls := getLinks(resp.Body)
	resp.Body.Close()

	urlMaps[url] = urls

	fmt.Printf("%*v found: %s with %d links\n", 4 * (maxDepth - depth), "", url, len(urls))

	for _, u := range urls {
		if strings.HasPrefix(u, "//") {
			u = "http:" + u
		}

		if strings.HasPrefix(u, "/") {
			u = baseUrl + u
		}

		if urlRegex.MatchString(u) {

			_, urlExists := urlMaps[u]
			if urlExists == false {
				fmt.Printf("%*v Exploring %s\n", 4 * (maxDepth - depth), "", u)
				Crawl(u, depth-1)
			} else {
				fmt.Printf("%*v Saved %s\n", 4 * (maxDepth - depth), "", u)
			}
		}
	}
	return
}

var maxDepth int

func main() {
	startTime := time.Now()
	maxDepth = 5
	urlMaps = make(map[string][]string)
	Crawl("http://motherfuckingwebsite.com/", maxDepth)

	now := time.Now()
	fmt.Printf("TOTAL TIME: %f\n", now.Sub(startTime).Seconds())
}