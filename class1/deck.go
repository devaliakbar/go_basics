package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func createCards() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" for "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {

	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func cardsFromFile(fileName string) deck {
	st, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(st), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().Local().UnixNano())
	r := rand.New(source)

	for i := range d {
		randomPosition := r.Intn(len(d) - 1)

		d[i], d[randomPosition] = d[randomPosition], d[i]
	}
}
