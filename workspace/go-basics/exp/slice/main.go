package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = [] string {"Red", "Blue","Green"}

	fmt.Println(colors[1:])

	ints:= make([]int, 5, 5)
	fmt.Println("initial capacity :", cap(ints))

	ints[0] = 33
	ints[1] = 22
	ints[2] = 55
	ints[3] = 122
	ints[4] = 1
	fmt.Println("capacity adding 5 elements :", cap(ints))
	ints = append(ints, 100)
	fmt.Println("adding 1 more element :", cap(ints))

	sort.Ints(ints)
	fmt.Println("In ascendeng order:", ints)

	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println("In ascendeng order:",  ints)

}
