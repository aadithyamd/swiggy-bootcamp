package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Connection Close")

	defer fmt.Println("Do the work 2")
	defer fmt.Println("Do the work 1")
	fmt.Println("Connection Open")

	//os.Exit(1)
	//panic("some Error")

	fmt.Println("Do the work")
	//fmt.Println("Do the work 1")
	//fmt.Println("Do the work 2")


}
