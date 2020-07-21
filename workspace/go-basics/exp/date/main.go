package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Date(2020,07,21,12,10,5,0,time.UTC)

	tomorrow := today.AddDate(0,0,1)

	fmt.Println(today.Format(time.RFC822Z))
	fmt.Println(tomorrow.Format(time.RFC822Z))
}
