package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://flipkart.com",
		"https://google.com",
		"https://youtube.com",
		"https://linkedins.com",
		"https://facebook.com",
	}

	c := make(chan string)

	for _, item := range links {
		go makeGetRequest(item, c)

	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			makeGetRequest(link, c)
		}(l)
	}
}

func makeGetRequest(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url + " is down!")
		c <- url
	} else {
		fmt.Println(url + " is up!")
		c <- url
	}

}
