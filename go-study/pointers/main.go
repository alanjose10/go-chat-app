package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	return &person{name: name, age: age}
}

func main() {

	person1 := newPerson("alan", 12)
	person2 := person1
	person1.name = "abc"

	fmt.Println(person1.name)
	fmt.Println(person2.name)

}
