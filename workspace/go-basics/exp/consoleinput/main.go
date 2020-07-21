package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//var s string
	//fmt.Scanln(&s)
	//fmt.Println(s)

	reader :=bufio.NewReader(os.Stdin)
	fmt.Println("Salary:")
	s, _ := reader.ReadString('\n')
	sal, _ := strconv.ParseFloat(strings.TrimSpace(s), 64);
	fmt.Println("Months:")
	m, _ := reader.ReadString('\n')
	months, _ := strconv.ParseInt(strings.TrimSpace(m),10, 64)

	fmt.Println("Salary into months = ", sal * float64(months))
}
