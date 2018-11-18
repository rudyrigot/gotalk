package gotalk

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func (p *person) speak() string {
	return fmt.Sprintf("Hello, I'm %s!", p.name)
}

// Memory does things about state management
func Memory() {
	david := person{"David", 53}

	// pointerToDavid := &david
	var pointerToDavid *person
	pointerToDavid = &david

	fmt.Println(david.speak())
	fmt.Println(pointerToDavid.speak())
}
