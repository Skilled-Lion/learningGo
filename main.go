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
	var cccc map[string]string
	bbbb := make(map[string]string)
	colors := map[string]string{
		"red":  "reeeed",
		"blue": "bbbbblllluuueee",
	}
	bbbb["aaaa"] = "aaas"
	printMap(colors)
	delete(colors, "red")
	fmt.Println(cccc, colors, bbbb)
}

func printMap(m map[string]string) {
	for a, b := range m {
		fmt.Println(a + " => " + b)
	}
}

func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
