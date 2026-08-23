package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func handle_error(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
		"http://golang.org",
	}

	// create a channel to receive messages from spawned go routines
	c := make(chan string)

	// create routines for each link
	for _, url := range links {
		go checkLink(url, c)
	}

	// await messages from the channel and dispatch checkLink again
	for l := range c {
		// function literal to inject pause before calling checkLing again
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkLink(url, c)
		}(l)
	}
}

func checkLink(url string, c chan string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(url, err)
		c <- url
		return
	}
	fmt.Println(url, resp.Status)
	c <- url
}
