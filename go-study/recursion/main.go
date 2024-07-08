package main

import (
	"fmt"
	"time"
)

func factorial(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var result uint64
	startTime := time.Now()
	result = factorial(4)
	endTime := time.Now()
	delta := endTime.Sub(startTime)
	fmt.Println(result)
	fmt.Println(delta.Seconds())
}
