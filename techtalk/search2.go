// +build OMIT

package main

import (
	"fmt"
	"net/http"
	"time"
)

// GSR START OMIT
func getSearchResult(search string) result {
	time_start := time.Now()
	res, _ := http.Get(fmt.Sprintf("http://search.businessoffashion.com/api/search?q=%s", search))
	time_dur := time.Since(time_start)
	return result{res, time_dur}
}

// GSR END OMIT

type result struct {
	res  *http.Response
	time time.Duration
}

// START1 START OMIT
func main() {
	time_start := time.Now()
	c := fanIn(search("armani"), search("hugo boss"), search("burberry")) // HL
	for i := 0; i < 3; i++ {
		val := <-c
		fmt.Println(val.time)
	}
	time_dur := time.Since(time_start)
	fmt.Println("---")
	fmt.Println(time_dur)
}

// START1 END OMIT

// START2 OMIT
func search(msg string) <-chan result { // Returns receive-only channel of strings. // HL
	c := make(chan result)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- getSearchResult(msg)
		}
	}()
	return c // Return the channel to the caller. // HL
}

// STOP2 OMIT

// START3 OMIT
func fanIn(input1, input2, input3 <-chan result) <-chan result { // HL
	c := make(chan result)
	go func() {
		for {
			c <- <-input1
		}
	}() // HL
	go func() {
		for {
			c <- <-input2
		}
	}() // HL

	go func() {
		for {
			c <- <-input3
		}
	}() // HL
	return c
}

// STOP3 OMIT
