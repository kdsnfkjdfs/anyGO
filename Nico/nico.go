package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {

		go hitURL(url, c)

	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, stauts := range results {
		fmt.Println(url, stauts)
	}

}

func hitURL(url string, c chan<- result) {
	fmt.Println("Hit URL: ", url)

	rep, err := http.Get(url)
	stare := "OK"

	if err != nil || rep.StatusCode >= 400 {
		stare = "Filed"
	}
	c <- result{url: url, status: stare}

}
