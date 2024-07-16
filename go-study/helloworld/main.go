package main

import (
	"fmt"
)

type A struct {
}

func (a A) TestMethod() {
	fmt.Println("TestMethod for A")
}

type B struct {
	a A
}

func (b B) TestMethod() {
	fmt.Println("TestMethod for B")
}

func main() {

	// a := A{}
	b := B{A{}}
	b.TestMethod()
	b.a.TestMethod()
}
