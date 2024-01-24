package main

import "fmt"

// this new deck type kinda of "extends" the []string slice built-in type
type deck []string

// receiver function, which allows to setup methods on user-defined types
// "define print() for all variables of type 'deck'"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
