package main

import (
	"fmt"
	"net/http"
	"time"
)

type Result struct {
	id  int
	url string
	res *http.Response
	err error
}

func SyncGets(urls []string) []*Result {
	results := []*Result{}
	for id, url := range urls {
		res, err := http.Get(url)
		if err == nil {
			res.Body.Close()
		}
		results = append(results, &Result{id, url, res, err})
	}
	return results
}
func AsyncGets(urls []string) []*Result {
	results := []*Result{}
	queue := make(chan *Result, len(urls))
	for id, url := range urls {
		go func(id int, url string) {
			res, err := http.Get(url)
			if err == nil {
				res.Body.Close()
			}
			queue <- &Result{id, url, res, err}
		}(id, url)
	}

	for {
		select {
		case r := <-queue:
			results = append(results, r)
			if len(results) == len(urls) {
				return results
			}
		case <-time.After(10 * time.Millisecond):
			fmt.Printf(".")
		}
	}
}

func MakeUrls(size int) []string {
	urls := make([]string, size)
	for i := 0; i < size; i++ {
		urls[i] = "http://localhost:3000/fib/10"
	}
	fmt.Println("Size urls", len(urls))
	return urls
}
func main() {
	urls := MakeUrls(100)
	start := time.Now()
	results := AsyncGets(urls)
	fmt.Printf("\nFetch %d urls in %v\n", len(urls), time.Since(start))
	for _, result := range results {
		if result.err != nil {
			fmt.Println("Error for ", result.url)
		} else {
			fmt.Println(result.id, result.url, result.res.Status)
		}
	}
}
