package main

import (
	"fmt"
	"runtime"
)

func channelTest() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <-messages
	fmt.Println(message1)

	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}

var Message = make(chan string)

func channelTestAgain() {
	runtime.GOMAXPROCS(2)

	go cetakAgain("Budi")
	data := <-Message
	fmt.Println(data)
}

func cetak(name string) string {
	return "hello" + name
}

func cetakAgain(name string) {
	data := "Hello " + name
	Message <- data
}
