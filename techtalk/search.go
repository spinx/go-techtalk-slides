// +build OMIT

package main

import (
	"fmt"
	"net/http"
	"time"
)

type result struct {
	res  *http.Response
	time time.Duration
}

// f START OMIT
func search(query string, ch chan result) {
	res, time := getSearchResult(query)

	send := result{res, time}

	ch <- send
}

// f END OMIT

// GSR START OMIT
func getSearchResult(q string) (*http.Response, time.Duration) {
	time_start := time.Now()
	res, _ := http.Get(fmt.Sprintf("http://search.businessoffashion.com/api/search?q=%s", q))
	time_dur := time.Since(time_start)
	return res, time_dur
}

// GSR END OMIT

type searchResult struct {
	title string
}

// START1 START OMIT
func main() {
	time_start := time.Now()

	getSearchResult("armani")
	getSearchResult("hugo boss")
	getSearchResult("burberry")

	time_dur := time.Since(time_start)
	fmt.Println(time_dur)
}

// START1 END OMIT
