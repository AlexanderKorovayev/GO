package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	size := make(chan int)
	go responseSize("https://example.com/", size)
	go responseSize("https://golang.org/", size)
	go responseSize("https://golang.org/doc", size)
	fmt.Println(<-size)
	fmt.Println(<-size)
	fmt.Println(<-size)
}

func responseSize(url string, channel chan int) {
	fmt.Println("getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}
