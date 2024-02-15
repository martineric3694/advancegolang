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

func stringReplace(data string, old string, new string, index int) {
	fmt.Println(strings.Replace(data, old, new, index))
}

func strPadLeft(data string) {
	fmt.Printf("%-8s", data)
}
