package main

import (
	"fmt"
	"strings"
)

func substring(data string, start int, end int) {
	fmt.Println("Hasil Substring : " + data[start:end])
}

func splitString(data string) {
	fmt.Println(strings.Split(data, " "))
}
