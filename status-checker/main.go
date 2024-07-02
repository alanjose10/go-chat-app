package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://google.com/alan",
		"https://youtube.com",
		"https://linkedins.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
		"https://twitter.com",
		"https://instagram.com",
		"https://reddit.com",
		"https://ebay.com",
		"https://wikipedia.org",
		"https://craigslist.org",
		"https://bing.com",
		"https://pinterest.com",
		"https://github.com",
		"https://wordpress.com",
		"https://apple.com",
		"https://tumblr.com",
		"https://microsoft.com",
		"https://netflix.com",
		"https://paypal.com",
		"https://imdb.com",
		"https://espn.com",
		"https://nytimes.com",
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.baidu.com",
		"https://www.yahoo.com",
		"https://flipkart.com",
		"https://amazon.in",
		"https://myntra.com",
	}

	c := make(chan string)

	for _, item := range links {
		go makeGetRequest(item, c)

	}
	fmt.Println(<-c)
}

func makeGetRequest(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		c <- url + " is down"
	} else {
		c <- url + " is up"
	}
}
