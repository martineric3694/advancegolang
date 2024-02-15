package main

import (
	"fmt"
	"runtime"
	"sync"
)

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func printWG(till int, message string, wg *sync.WaitGroup) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
	wg.Done()
}

func waitGroup() {
	runtime.GOMAXPROCS(2)

	// Ada 3 key dalam WaitGroup : Add / Wait / Done
	var wg sync.WaitGroup

	wg.Add(1)
	go printWG(100, "Hallo", &wg)

	wg.Add(1)
	go printWG(30, "Dunia", &wg)

	wg.Wait()
}
