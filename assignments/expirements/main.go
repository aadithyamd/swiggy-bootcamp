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
	fmt.Printf("EmpID : %v \n", e.empid)
	fmt.Printf("Name : %v \n", e.empname)

	fmt.Printf("Salary : %v \n", e.salary)
	fmt.Printf("EmpID : %v \n", e.empid)
	fmt.Printf("Address : %v \n", e.address)

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
	fmt.Println("Enter the details: \n")
	enterMore := true
	i := 0
	for enterMore == true {
		enterMore = false
		fmt.Println("Employee ID:")
		eid, _ := reader.ReadString('\n')
		employee.empid = eid
		fmt.Printf("%v \n", eid)

		fmt.Println("Employee Name:")
		name, _ := reader.ReadString('\n')
		employee.empname = name
		fmt.Printf("%v \n", name)

		fmt.Println("Salary:")
		strSal, _ := reader.ReadString('\n')
		sal, _ := strconv.ParseFloat(strings.TrimSpace(strSal),64)
		employee.salary = sal
		fmt.Printf("%v \n", sal)


		fmt.Println("Address :\n")

		fmt.Println("House no :")
		hno, _ := reader.ReadString('\n')
		employee.address.hno = hno
		fmt.Printf("%v \n", hno)


		fmt.Println("Street :")
		street, _ := reader.ReadString('\n')
		employee.address.street = street
		fmt.Printf("%v \n", street)

		fmt.Println("City :")
		city, _ := reader.ReadString('\n')
		employee.address.city = city
		fmt.Printf("%v \n", city)

		fmt.Println("Pin :")
		pin, _ := reader.ReadString('\n')
		employee.address.pin = pin
		fmt.Printf("%v \n", pin)

		if (i % 5 == 0 && i > 0) {
			emps = append(emps, employee)
		} else {
			emps[i] = employee
		}

		fmt.Println("Do you want to continue? (N to quit any other key to continue ..)")
		cont, _ := reader.ReadString('\n')
		fmt.Printf("%v \n", cont)

		if strings.TrimSpace(cont) == "N" {
			enterMore = false
		} else {
			enterMore = true
		}

		i++
	}

	employees := make([] Animal, i) //[] Animal{emps}
	for i, e := range emps {
		employees[i] = e

		fmt.Printf("\nEmployee %v:\n", i)
		_ = employees[i].Speak()
	}

	//strs := []string{"first", "second"}
	//names := make([]interface{}, len(strs))
	//for i, s := range strs {
	//	names[i] = s
	//}

	fmt.Println(employees)
	//for
}
