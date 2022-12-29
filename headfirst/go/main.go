package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url  string
	Size int
}

func main() {
	page := make(chan Page)
	go responseSize("http://www.baidu.com", page)
	go responseSize("http://www.mi.com", page)

	channel1 := make(chan string)
	channel2 := make(chan string)
	go abc(channel1)
	go def(channel2)

	fmt.Println(<-page)
	fmt.Println(<-page)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
}

func abc(channel chan string) {
	channel <- "a"
	channel <- "b"
	channel <- "c"
}

func def(channel chan string) {
	channel <- "d"
	channel <- "e"
	channel <- "f"
}

func responseSize(url string, page chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	page <- Page{Url: url, Size: len(string(body))}
}
