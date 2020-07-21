package main

import "fmt"

func main() {
	i1, i2, i3 := 10, 20, 30

	intSum := i1 + i2 + i3
	fmt.Println("Sum of Integers", intSum)

	f1, f2, f3 := 1.1, 2.2, 3.3

	fSum := f1 + f2 + f3

	fmt.Println("Sum of floats", fSum)
}
