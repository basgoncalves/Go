package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // not need the name of
}

// two different ways to use structure
func main() {

	// write struct version 1
	// alex := person{firstName: "Alex", lastName: "Andrerson"}
	// fmt.Println(alex)

	// version 2
	var alex person
	alex.lastName = "Anderson"
	alex.firstName = "Alex"
	alex.contact.email = "AA@gmail.com"
	alex.contact.zipCode = 12345

	// update struct version 1
	alex = alex.updateName("Alexander")
	alex.print()

	// update struct version 2 (with a pointer)
	alexPointer := &alex
	alexPointer.updateNamePointer("Alexander")
	alexPointer.print()

	// slices work different than structs example
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)

}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p person) updateName(newFristName string) person {
	p.firstName = newFristName
	return p
}

func (pointerToPerson *person) updateNamePointer(newFristName string) {
	(*pointerToPerson).firstName = newFristName
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
