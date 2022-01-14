package main

func main() {
	cards := cardsFromFile("jose")
	cards.shuffle()
	cards, _ = deal(cards, 3)
	cards.print()
}
