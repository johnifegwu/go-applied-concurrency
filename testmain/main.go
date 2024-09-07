package server

import (
	"fmt"
	"sync"
)

func testmain() {
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()
	goodbye()
}

func hello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello world")
}

func goodbye() {
	fmt.Println("Goodbye world")
}
