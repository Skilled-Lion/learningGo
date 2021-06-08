package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "jim",
		lastName:  "jam",
		contactInfo: contactInfo{
			email:   "jim@j.com",
			zipCode: 182839,
		},
	}

	jim.updateName("jiiim")
	jim.print()
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
