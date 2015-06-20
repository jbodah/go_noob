package main

import (
	"fmt"
)

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func NewSaiyan(name string, power int, father *Saiyan) *Saiyan {
	return &Saiyan{name, power, father}
}

func main() {
	goku := NewSaiyan("Goku", 9001, nil)
	gohan := NewSaiyan("Gohan", 1000, goku)
	fmt.Println(gohan.Father.Name)
}
