package main


func main() {
	// create a deck and save it
	// cards := newDeck()

	// hand, _ := deal(cards, 5)

	// hand.print()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my cards")

	//load deck from file
	cards := newDeckFromFile("my cards")
	cards.shuffle()
	cards.print()

}
