package main

import (
	"fmt"
)

type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = &Saiyan{"Krillin", nil}
	fmt.Println(goku.Friends["krillin"].Name)

	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
	}

	// map iteration is unordered
	for key, value := range lookup {
		fmt.Println(key, value)
	}
}
