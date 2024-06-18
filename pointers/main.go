package main

import "fmt"

func main() {

	mySlice := []string{"Hi", "How", "Are", "You"}

	fmt.Println(mySlice)

	updateSlice(mySlice)

	fmt.Println(mySlice)

}

func updateSlice(s []string) {
	s[0] = "Bye"
}
