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
	type urlCheck struct {
		url       string
		processed chan bool
	}

	var urlChecker chan *urlCheck
	var isRoot bool

	for _, internalArg := range internalArgs {
		switch argVal := internalArg.(type) {
		case chan *urlCheck:
			urlChecker = argVal
		default:
			panic(fmt.Sprintf("unknown argument type: %T", internalArg))
		}
	}

	if urlChecker == nil {
		isRoot = true
		urlChecker = make(chan *urlCheck)
		go func() {
			fmt.Println("checker started")
			visitedUrls := make(map[string]bool)
			for {
				check := <-urlChecker
				if check.url == "" {
					break
				}
				if _, ok := visitedUrls[check.url]; ok {
					check.processed <- true
				} else {
					visitedUrls[check.url] = true
					check.processed <- false
				}
			}
			fmt.Println("checker stopped")
		}()
	}

	if depth <= 0 {
		return
	}

	var body string
	var urls []string
	var err error

	check := &urlCheck{url: url, processed:make(chan bool)}
	urlChecker <- check
	if (<-check.processed) {
		return
	}

	body, urls, err = fetcher.Fetch(url)
	if err != nil {
		ch <- err.Error()
		return
	}

	ch <- fmt.Sprintf("found: %s %q", url, body)
	innerWaiting := new(sync.WaitGroup)
	innerWaiting.Add(len(urls))
	for _, u := range urls {
		go func(url string) {
			defer innerWaiting.Done()
			Crawl(url, depth-1, fetcher, ch, urlChecker)
		}(u)
	}
	innerWaiting.Wait()
	if isRoot {
		urlChecker <- &urlCheck{}
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