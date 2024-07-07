package main

import "fmt"

type address struct {
	houseNumber string
	streetName  string
	postCode    string
	country     string
}

type person struct {
	firstName string
	lastName  string
	age       int
	address
}

func (p person) printDetails() {
	fmt.Println("Person details")
	fmt.Printf("%+v\n", p)
}

func (personPointer *person) setFirstName(firstName string) {
	(*personPointer).firstName = firstName
}

func main() {

	alan := person{
		firstName: "Alan",
		lastName:  "Jose",
		age:       29,
		address: address{
			houseNumber: "test",
			streetName:  "test street",
			postCode:    "1234ABC",
			country:     "United Kingdom",
		},
	}

	alan.printDetails()

	alanPointer := &alan
	alan.setFirstName("Baby")
	alan.printDetails()

	fmt.Println(&alanPointer)

}
