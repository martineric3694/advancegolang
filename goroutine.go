package main

import (
	"fmt"
	"runtime"
)

func implGoroutine() {
	// Goroutine
	// Async proses, dimana print biasa lebih lama diprosesnya daripada menggunakan go print
	runtime.GOMAXPROCS(2)

	// // asynchronous
	go print(5, "halo")
	go print(10, "selamat")
	go print(2, "datang")

	// // synchronous
	print(5, "apa kabar - tanpa goroutine")
	print(10, "selamat - tanpa goroutine")
	print(10, "datang - tanpa goroutine")

	var input string
	fmt.Scanln(&input)

	waitGroup()
}
