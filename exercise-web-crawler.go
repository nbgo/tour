package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan string, internalArgs ...interface{}) {
	var terminator chan bool
	var visitedUrls struct { sync.RWMutex; m map[string]bool }
	for _, internalArg := range internalArgs {
		switch argVal := internalArg.(type) {
		case chan bool:
			terminator = argVal
			defer func() {terminator <- true}()
		case struct { sync.RWMutex; m map[string]bool }:
			visitedUrls = argVal
		default:
			panic(fmt.Sprintf("Unknown argument type: %T", internalArg))
		}
	}

	if visitedUrls.m == nil {
		visitedUrls = struct{
			sync.RWMutex
			m map[string]bool
		}{m: make(map[string]bool)}
	}

	if depth <= 0 {
		return
	}

	if func() bool {
		if func() bool {
			visitedUrls.RLock()
			defer visitedUrls.RUnlock()
			_, ok := visitedUrls.m[url]
			if ok {
				return true
			}
			return false
		}() {
			return true
		}

		visitedUrls.Lock()
		defer visitedUrls.Unlock()
		// We need to check again because between RUnlock and Lock changes can happen
		_, ok := visitedUrls.m[url]
		if ok {
			return true
		}
		visitedUrls.m[url] = true
		return false
	}() {
		return
	}


	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		ch <- err.Error()
		return
	}
	ch <- fmt.Sprintf("found: %s %q", url, body)
	innerCrawlDone := make(chan bool)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, ch, innerCrawlDone, visitedUrls)
	}
	for range urls {
		<-innerCrawlDone
	}
	if terminator == nil {
		close(ch)
	}
	return
}

func main() {
	ch := make(chan string)
	go Crawl("http://golang.org/", 4, fetcher, ch)
	for url := range ch {
		fmt.Println(url)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}