package main

import "fmt"

type Animal interface {
	// method name return type
	Speak() string
}

type Dog struct {}
type Cat struct {}
type Emp struct {}

func (e Emp) Speak() string {
	return "Hello.. Hello"
}
func (d Dog) Speak() string {
	return "Mmm.. Wof"
}
func (c Cat) Speak() string {
	return "Mew.. Mew"
}

func  main()  {
	dog1 := Dog{}
	fmt.Println(dog1)
	fmt.Println(dog1.Speak())

	animals := []Animal {Dog{}, Emp{}, Cat{}}
	for e, a := range animals{
		fmt.Println(a.Speak())
		fmt.Println(e)
	}
}
