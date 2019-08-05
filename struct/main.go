package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func (p *person) updateFirstName(name string) {
	(*p).firstname = name
}

func main() {
	p := person{"John", "Smith"}
	fmt.Println(p.firstname)

	k := &p
	k.updateFirstName("updated")
	fmt.Println(k.firstname)
}
