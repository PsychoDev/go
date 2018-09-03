package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := make(chan string)
	sites := []string{
		"http://pimcore.com",
		"http://stackoverflow.com",
		"http://google.com",
		"http://golang.org",
	}

	for _, link := range sites {
		go fetchLink(link, c)
	}

	<-c
	<-c
	<-c
	<-c
}

func fetchLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println("Fetched:", link)
	c <- link
}
