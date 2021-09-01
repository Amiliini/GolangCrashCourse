package main

import (
	"fmt"
	"strconv"
)

//define person struct

type Person struct {
	//	firstName string
	//	lastName  string
	//	city      string
	//	gender    string
	//	age       int

	firstName, lastName, city, gender string
	age                               int
}

//Greeting method(value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

//hasBirthday method (poitner receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)

func (p *Person) getMarried(spouceLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouceLastName
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Jaska", lastName: "Jokunen", city: "Hki", gender: "m", age: 30}

	person2 := Person{"Sylvi", "Kissanen", "TKU", "f", 26}
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person2.city)
	person1.age++
	person1.hasBirthday()
	fmt.Println(person1.greet())
	person2.getMarried("Williams")
	fmt.Println(person2.greet())
	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
}
