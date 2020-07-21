package main

import "fmt"

func main() {

	str1 := "This is string 1"
	aNumber := 33
	isTrue := true
	strlen, err := fmt.Println(str1);

	if err == nil {
		fmt.Println(strlen)
	}

	fmt.Printf("value of num as float %.2f", float64(aNumber))

	fmt.Printf("Data types : %T %T %T", str1, aNumber,isTrue)
}