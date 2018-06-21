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
 

type helper struct {
	cache map[string]bool
	m sync.Mutex
	wg sync.WaitGroup
}


func (h *helper) Crawl(url string, depth int, fetcher Fetcher) {
	h.wg.Add(1)
	h._crawl(url, depth, fetcher)
	h.wg.Wait()
}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (h *helper) _crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
		
	defer h.wg.Done()
	
	if depth <= 0 {
		return
	}
	
	h.m.Lock()
	if h.cache[url] {
		h.m.Unlock()
		return
	}
	
	h.cache[url] = true
	h.m.Unlock()
	
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
		
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {	
		h.wg.Add(1)
		go h._crawl(u, depth-1, fetcher)	
	}	
	return
}

func main() {
	h := &helper{cache: make(map[string]bool)}	
	h.Crawl("https://golang.org/", 4, fetcher)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
