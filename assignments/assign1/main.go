package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface {
	Eat() string
	Speak() string
	Breath() string
}

func (e Emp) Speak() string{
	fmt.Printf("EmpID : %v", e.empid)
	fmt.Printf("Name : %v", e.empname)

	fmt.Printf("Salary : %v\n", e.salary)
	fmt.Printf("EmpID : %v", e.empid)
	fmt.Printf("Address  \n") //hno, street, city, pin
	fmt.Printf("House No : %v", e.address.hno)
	fmt.Printf("Street : %v", e.address.street)
	fmt.Printf("City : %v", e.address.city)
	fmt.Printf("PIN : %v\n", e.address.pin)

	return ""
}

func (e Emp) Breath() string{
	return "haaaahh...."
}

func (e Emp) Eat() string{
	return "slurp...slurp....."
}





type Address struct {
	hno, street, city, pin string
}
type Emp struct {
	empid, empname string
	salary float64
	address Address
}

func main(){

	var employee Emp
	//	var eid, name, strSal string

	reader := bufio.NewReader(os.Stdin)
	emps := make([]Emp,5,5)
	fmt.Printf("Enter the details: \n")
	enterMore := true
	i := 0
	for enterMore == true {
		enterMore = false
		fmt.Printf("Employee ID:")
		eid, _ := reader.ReadString('\n')
		employee.empid = eid
		//fmt.Printf("%v \n", eid)

		fmt.Printf("Employee Name:")
		name, _ := reader.ReadString('\n')
		employee.empname = name
		//fmt.Printf("%v \n", name)

		fmt.Printf("Salary:")
		strSal, _ := reader.ReadString('\n')
		sal, _ := strconv.ParseFloat(strings.TrimSpace(strSal),64)
		employee.salary = sal
		//fmt.Printf("%v \n", sal)


		fmt.Printf("Address :\n")

		fmt.Printf("House no :")
		hno, _ := reader.ReadString('\n')
		employee.address.hno = hno
		//fmt.Printf("%v \n", hno)


		fmt.Printf("Street :")
		street, _ := reader.ReadString('\n')
		employee.address.street = street
		//fmt.Printf("%v \n", street)

		fmt.Printf("City :")
		city, _ := reader.ReadString('\n')
		employee.address.city = city
		//fmt.Printf("%v \n", city)

		fmt.Printf("Pin :")
		pin, _ := reader.ReadString('\n')
		employee.address.pin = pin
		//fmt.Printf("%v \n", pin)

		if (i % 5 == 0 && i > 0) {
			emps = append(emps, employee)
		} else {
			emps[i] = employee
		}

		fmt.Printf("Do you want to continue? (N to quit any other key to continue ..) :")
		cont, _ := reader.ReadString('\n')
		//fmt.Printf("%v \n", cont)

		if strings.TrimSpace(cont) == "N" || strings.TrimSpace(cont) == "n" {
			enterMore = false
		} else {
			enterMore = true
		}

		i++
	}

	employees := make([] Animal, i)
	for in, e := range emps {
		if(in >= i) {
			break
		}
		employees[in] = e

		fmt.Printf("\nEmployee %v\n\n", in + 1)
		_ = employees[in].Speak()
	}

}
