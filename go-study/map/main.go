package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	fmt.Println(colors)

	printMap(colors)

	fmt.Println(colors)

}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Printf("Hex code of %s is %s\n", color, hex)
	}
	delete(m, "black")
}
