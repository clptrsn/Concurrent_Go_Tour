package main

import (
	"fmt"
	"net/http"
	"io"
	"golang.org/x/net/html"
	"regexp"
	"strings"
	"time"
	"errors"
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

type result struct {
	baseUrl string
	url string
	urls []string
	depth int
	err error
}

type urlMapInner struct {
	finished bool
	urlChildren []string
}

var urlMaps map[string]urlMapInner

func Crawl(url string, depth int) {
	urlRegex, _ := regexp.Compile(`(?P<Scheme>http(?:s|):\/\/)(?P<Path>.*)(?:\/|$)`)
	urlResultChan := make(chan result)

	// Function which will fetch the child urls. It will pass the result into the result channel
	fetch := func(url string, depth int) {
		if depth <= 0 {
			// send a result with a depth of 0
			urlResultChan <- result{depth:0}
			return
		}

		// checks if the url matches the regex above
		isValidUrl := urlRegex.MatchString(url)
		if isValidUrl == false {
			// send a result with an error
			urlResultChan <- result{err: errors.New("Url Wrong Format") }
			return
		}

		// Extract the protocol and base url path
		match := urlRegex.FindStringSubmatch(url)
		baseUrl := match[1] + match[2]

		// Get the webpage from the url
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			urlResultChan <- result{err: err}
			return
		}
		// Get the <a> element's href property from the page
		urls := getLinks(resp.Body)
		resp.Body.Close()

		fmt.Printf("%*v found: %s with %d links\n", 4 * (maxDepth - depth), "", url, len(urls))
		
		// Add the url to the global map
		urlMaps[url] = urlMapInner{finished:true, urlChildren: urls}
		
		// Send the result to the channel
		urlResultChan <- result{baseUrl: baseUrl, url: url, urls: urls, depth: depth}
	}

	urlMaps[url] = urlMapInner{finished: false, urlChildren: []string{}}
	go fetch(url, depth)

	// For each url we find, we will increase urlsProcessing
	for urlsProcessing := 1; urlsProcessing > 0; urlsProcessing-- {
		// Get a result from the results channel
		result := <-urlResultChan

		if result.err != nil {
			continue
		}

		if result.depth <= 0 {
			continue
		}

		for _, u := range result.urls {
			if strings.HasPrefix(u, "//") {
				u = "http:" + u
			}

			if strings.HasPrefix(u, "/") {
				u = result.baseUrl + u
			}

			if urlRegex.MatchString(u) {

				_, urlExists := urlMaps[u]
				if urlExists == false {
					fmt.Printf("%*v Exploring %s\n", 4 * (maxDepth - result.depth), "", u)
					
					urlsProcessing++
					urlMaps[u] = urlMapInner{finished: false, urlChildren: []string{}}
					go fetch(u, result.depth-1)
				} else {
					fmt.Printf("%*v Saved %s\n", 4 * (maxDepth - result.depth), "", u)
				}
			}
		}
	}

	close(urlResultChan)
	return
}

var maxDepth int

func main() {
	startTime := time.Now()

	maxDepth = 5
	urlMaps = make(map[string]urlMapInner)
	Crawl("http://motherfuckingwebsite.com/", maxDepth)

	now := time.Now()
	fmt.Printf("TOTAL TIME: %f\n", now.Sub(startTime).Seconds())
}