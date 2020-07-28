package main

import "fmt"

func main(){
	fmt.Printf("Name is %v, Length is %v",
		AddAllValues(1))
}


func AddAllValues(value ... int) (string, int) {
	sum := 0
	for _,v := range value{
		sum += v
	}
	name := "Harry"
	return name, sum
}
