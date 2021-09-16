package main

import (
	"fmt"
	"net/http"
	"io"
	"regexp"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}


type WebCrawlerError struct {	
	Message string
}

type WebCrawler struct{
	visited map[string]int
	mu sync.Mutex
}

func (e WebCrawlerError) Error() string {
	return e.Message
}

type MyFetcher struct{
	name string
}



func (fetcher MyFetcher) Fetch(url string) (body string, urls []string, err error){
	
	
	res, err := http.Get(url)
	if err != nil {
		return "", make([]string, 0), err
	}


	bodyStr, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return "", make([]string, 0), WebCrawlerError{fmt.Sprintf("Response failed with status code :at %v", res.StatusCode)}
	}
	if err != nil {
		return "", make([]string,0 ), err
	}
	
	re := regexp.MustCompile(`(http|ftp|https)://([\w+?\.\w+])+([a-zA-Z0-9\~\!\@\#\$\%\^\&\*\(\)_\-\=\+\\\/\?\.\:\;\'\,]*)?`)
	result := re.FindAll(bodyStr, -1)
	re_urls := make([]string, len(result))
	for i:=0; i<len(result); i++{

      
      re_urls[i] = string(result[i])
      
      
    }
	
	return string(bodyStr), re_urls, nil
	
	
}

func (crawler *WebCrawler ) checkVisited(url string) bool{
	crawler.mu.Lock()
	defer crawler.mu.Unlock()
	return crawler.visited[url] > 0
}

func (crawler *WebCrawler ) markVisited(url string){
	crawler.mu.Lock()
	crawler.visited[url]++
	crawler.mu.Unlock()

}


// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (crawler *WebCrawler)Crawl(url string, depth int, fetcher Fetcher) {

	if depth <= 0 {
		return
	}

	if crawler.checkVisited(url){
		//fmt.Printf("visited %s\n", url)
		return
	}

	_, urls, err := fetcher.Fetch(url)

	crawler.markVisited(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nfound: %s\n", url)
	for _, u := range urls {
		
		crawler.Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	
	fetcher:= MyFetcher{name: "rachel's fetcher"}
	crawler:= &WebCrawler{visited: make(map[string]int)}
	for i := 0; i < 100; i++ {
		crawler.Crawl("https://golang.org/", 3, fetcher)
		
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
