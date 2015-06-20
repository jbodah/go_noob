package main

import (
	"fmt"
)

type Person struct {
	Name string
}

type Saiyan struct {
	// Composition/trait/mixin
	*Person
	Power int
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
}
