package main

func main() {

	fileName := "my_cards"

	cards := newDeck()

	cards.shuffleDeck()
	cards.printDeck()

	cards.saveToFile(fileName)

	// sl := []int{}

	// for i := range 10 {
	// 	sl = append(sl, i)
	// }

	// for _, v := range sl {
	// 	if v%2 == 1 {
	// 		fmt.Println("odd")
	// 	} else {
	// 		fmt.Println("even")
	// 	}
	// }
}
