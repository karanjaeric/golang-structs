package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	processAccount()
}

type Account struct {
	accountNumber string
	accountName   string
	balance       float64
	active        bool
}

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func processAccount() {
	// var customerAccount Account
	customerAccount := Account{}

	customerAccount.accountNumber = "123456789"
	customerAccount.accountName = "ERICK KARANJA"
	customerAccount.balance = 150000
	customerAccount.active = true
	fmt.Printf("Account details %v\n", customerAccount)

	//Anonymous structs
	var person struct {
		name string
		age  int
	}
	person.age = 10
	person.name = "EricK Karanja Muthike"
	fmt.Printf("Person is %v\n", person)

	pet := struct {
		name string
		kind string
	}{
		"Kido", "dog",
	}
	fmt.Printf("Pet is %v\n", pet)

	student := struct {
		name    string
		faculty string
	}{
		name:    "Sheikh Moha",
		faculty: "Engineering",
	}

	fmt.Printf("Student is %v\n", student)

	employeeOne := Employee{"Erick", "Karanja", 1}
	fmt.Printf("Employee 1 is %v\n", employeeOne)
	employeeTwo := Employee{
		firstName: "Ninja",
		lastName:  "Afisti",
		id:        2,
	}
	fmt.Printf("Employee 2 is %v\n", employeeTwo)

	var employeeThree Employee
	employeeThree.firstName = "SuperMan"
	employeeThree.lastName = "Muthike"
	employeeThree.id = 3

	fmt.Printf("Employee 3 is %v\n", employeeThree)

}
