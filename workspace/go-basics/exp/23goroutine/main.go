package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://golang.org/",
		"https://jsonplaceholder.typicode.com/posts",
		"https://httpbin.org/xml",
	}

	var wg sync.WaitGroup
	for _, url := range urls{
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)

		wg.Wait()
	}
}

func returnType(url string) {
	resp, err := http.Get(url)
	checkError(err)

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	fmt.Println(url, ctype)
}

func checkError(err error){
	if err != nil{
		panic(err)
	}
}
