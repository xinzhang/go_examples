package main

import "fmt"

type bot interface {
	getHello() string
}

type englishBot struct{}
type spanishBot struct{}

func (eb englishBot) getHello() string {
	return "Hey"
}

func (sb spanishBot) getHello() string {
	return "Halo"
}

func sayHello(b bot) {
	fmt.Println(b.getHello())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	sayHello(eb)
	sayHello(sb)
}
