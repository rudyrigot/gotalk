package gotalk

import (
	"fmt"
	"sync"
	"time"
)

// Concurrency2 shows a second example about concurrency
func Concurrency2() {
	var wg sync.WaitGroup
	wg.Add(4)

	go waitABitAgain(3, &wg)
	go waitABitAgain(2, &wg)
	go waitABitAgain(1, &wg)
	go waitABitAgain(0, &wg)

	wg.Wait()
}

func waitABitAgain(s int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(s) * time.Second)
	fmt.Printf("I just waited %ds!\n", s)
}
