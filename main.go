package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var kaushal person
	kaushal.firstName = "kaushal"
	kaushal.lastName = "singh"
	alex := person{firstName: "alex", lastName: "anderson"}
	fmt.Println(alex)
	fmt.Println(kaushal)
	fmt.Printf("%+v", alex)
}
