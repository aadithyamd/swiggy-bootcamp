package main

import "fmt"

type Dog struct {
	Breed string
	Weight int

}

func main(){
	tutu := Dog{Breed: "Local", Weight: 15}
	fmt.Printf("Breed: % v, Weight %v", tutu.Breed, tutu.Weight)
	fmt.Println(tutu)

}
