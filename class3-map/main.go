package main

import "fmt"

func main() {

	//var person = map[string]string{}

	// person := make(map[string]string)

	// person["name"] = "Ali"

	person := map[string]string{
		"name":  "Jose",
		"age":   "45",
		"color": "black",
	}

	printMap(person)

}

func printMap(p map[string]string) {
	for key, value := range p {
		fmt.Println("Key :", key, ", Value :", value)
	}
}
