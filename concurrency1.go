package gotalk

import (
	"fmt"
	"time"
)

// Concurrency1 shows a first example about concurrency
func Concurrency1() {
	waitABit(3)
	waitABit(2)
	waitABit(1)
	waitABit(0)
}

func waitABit(s int) {
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Printf("I just waited %ds!\n", s)
}
