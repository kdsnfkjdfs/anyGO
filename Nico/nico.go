package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

func main() {

	results := make(map[string]string)
	c := make(chan string)
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
		result := "Good"
		err := hitURL(url)

		if err != nil {
			result = "Baby_Cat"
		}

		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, ": ", result)
	}

	go COUNT("BABY_CAT", c)
	go COUNT("NIGKRMAN", c)

	fmt.Println(<-c)
	fmt.Println(<-c)
}

func hitURL(url string) error {
	fmt.Println("Hit URL: ", url)

	rep, err := http.Get(url)

	if err != nil || rep.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}

func COUNT(name string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- name + "cat"

}
