package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type  Posts struct {
	Title  string
	UserId int
}

func main ()  {
	url := "https://jsonplaceholder.typicode.com/posts"
	data := ContentFromServer(url)
	fmt.Println(data)
	fmt.Println(PostFromJson(data))
}
func PostFromJson(content string) [] Posts {
	posts := make([]Posts, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))

	_, err := decoder.Token()
	CheckError(err)

	var post Posts

	for decoder.More() {
		err := decoder.Decode(&post)
		CheckError(err)
		posts = append(posts, post)
	}

	return posts
}

func ContentFromServer(url string) string {
	resp, err := http.Get(url)
	CheckError(err)
	defer resp.Body.Close()

	bytes,err := ioutil.ReadAll(resp.Body)
	CheckError(err)

	return string(bytes)
}

func CheckError(err error){
	if err != nil {
		fmt.Println("Error Occured !!\n")
		panic(err)
	}

}