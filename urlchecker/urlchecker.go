package urlchecker

import (
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func URLchecker(urls []string) {

	var results = map[string]string{}
	c := make(chan requestResult)

	for _, url := range urls {
		go hitURL(url, c)
	}
	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string, c chan<- requestResult) {
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- requestResult{url: url, status: status}
}
