package main

import "fmt"

type contactinfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactinfo contactinfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Williams",
		contactinfo: contactinfo{
			email:   "jim@gmail.com",
			zipCode: 98401,
		},
	}
	jim.updateName("James")
	jim.print()
}
