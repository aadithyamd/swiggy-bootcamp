package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url:="https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)

	if (err != nil) {
		errorHandler(err)
	}
	defer resp.Body.Close()
	fmt.Println("Content type", resp)

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errorHandler(err)
	}

	content := string(bytes)
	fmt.Println(content)
}

func errorHandler(err error){
	fmt.Println("Error Occured !!\n")
	panic(err)
}