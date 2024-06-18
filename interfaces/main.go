package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printReading(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printReading(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printReading(eb)
	// printReading(sb)
}
