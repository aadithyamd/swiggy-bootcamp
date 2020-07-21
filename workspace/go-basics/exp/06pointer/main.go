package main

import "fmt"

func main() {
	var p *int
	if p != nil {
		fmt.Println("value of P: ", *p);
	} else {
		fmt.Println("value of P: is nil ");
	}

	var v int = 44
	p = &v
	if p != nil {
		fmt.Println("value of P: ", *p);
	} else {
		fmt.Println("value of P: is nil ");
	}

}
