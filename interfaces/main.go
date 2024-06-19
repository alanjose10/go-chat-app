package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printReading(b bot) {
	fmt.Println(b.getGreeting())
}

// func printReading(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printReading(eb)
	printReading(sb)
}
