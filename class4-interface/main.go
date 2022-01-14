package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type arabicBot struct{}

func main() {
	eb := englishBot{}
	ab := arabicBot{}

	printGreeting(eb)
	printGreeting(ab)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (arabicBot) getGreeting() string {
	return "Assalamu Alaikkum!"
}
