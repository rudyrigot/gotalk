package gotalk

import (
	"fmt"
)

// State does everything about state management
func State() {
	d := dog{name: "Spot", age: 3}
	fmt.Println(d.speak())
	makeItSpeak(d) // interface
	// embedded
	dalmatian := dalmatian{NbOfSpots: 23}
	dalmatian.dog.name = "Rod"
	fmt.Println(dalmatian.speak())
}

// ---- STRUCT ----

type dog struct {
	name string
	age  int
}

func (d dog) speak() string {
	return fmt.Sprintf("Woof, my name is %s!", d.name)
}

// ---- INTERFACE ----

type animal interface {
	speak() string
}

func makeItSpeak(a animal) {
	fmt.Println(a.speak())
}

// ---- EMBEDDED ----

type dalmatian struct {
	dog
	NbOfSpots int
}
