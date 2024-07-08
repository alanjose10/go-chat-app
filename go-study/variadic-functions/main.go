package main

import "fmt"

func sums(numbers ...int) int {
	var total int
	for _, n := range numbers {
		total = total + n
	}
	return total
}

func main() {

	fmt.Println(sums(1, 2, 3, 4, 5, 6, 7, 8))

	numbers := []int{10, 20, 30}

	fmt.Println(sums(numbers...))
}
