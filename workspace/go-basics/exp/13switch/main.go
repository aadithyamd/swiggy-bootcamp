package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	day := rand.Intn(6) + 1
	result := ""

	switch day{
	case 1:
		result = "Sun"

	case 2:
		result = "Sat"
	default:
		result = "week day"
	}

	fmt.Println(result)
}
