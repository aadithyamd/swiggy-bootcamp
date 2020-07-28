package main

import "fmt"

var isConnected bool = false

func connect(){
	isConnected = true
	fmt.Println("DB connection opened")
}

func disconnect(){
	isConnected = false
	fmt.Println("DB connection clossed")
}

func main() {

}

