package main

func main() {
	cards := newDeck()
	// fmt.Println(cards.toString())
	// fmt.Println(cards)
	// fmt.Println([]string(cards))
	// cards.saveToFile(("my_cards"))
	// newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()

}
