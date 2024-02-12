package main

import (
	"fmt"
	"sync"
)

// Case
// restoran -> checking belanjaan (ikan, daging, sayur, buah, bawang)
//
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur -> kupas -> simpan box -> simpan chiler
// buah -> cuci -> simpan chiler
// bawang -> kupas -> simpan box -> simpan chiler
// jika bukan barang diatas -> buang sampah

func wgRoutine() {

	var wg sync.WaitGroup

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bumbu", "buah", "sterofoam", "bawang", "ikan",
		"ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang", "buah",
		"sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang", "buah", "sterofoam", "ikan",
	}
	i := 1
	for _, brg := range belanja {
		fmt.Print(i)
		fmt.Print(" ")
		switch brg {
		case "ikan":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go cuciWg(brg, &wg)
			i++
		case "daging":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go potongWg(brg, &wg)
			i++
		case "sayur":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go kupasWg(brg, &wg)
			i++
		case "buah":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go cuciWg(brg, &wg)
			i++
		case "bawang":
			fmt.Println("Menemukan", brg)
			wg.Add(1)
			go kupasWg(brg, &wg)
			i++
		default:
			wg.Add(1)
			go buangSampahWg(brg, &wg)
			i++
		}
	}
	wg.Wait()
}

func cuciWg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Cuci", barang)
	if barang == "ikan" {
		simpanBoxWg(barang, wg)
	} else {
		simpanChilerWg(barang, wg)
	}
}

func potongWg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Potong", barang)
	simpanFrozenWg(barang, wg)
}

func kupasWg(barang string, wg *sync.WaitGroup) {
	fmt.Println("Cuci", barang)
	simpanBoxWg(barang, wg)
}

func simpanBoxWg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Box")
	simpanChilerWg(barang, wg)
}

func simpanFrozenWg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Frozen")
	wg.Done()
}

func simpanChilerWg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Simpan Chiler")
	wg.Done()
}

func buangSampahWg(barang string, wg *sync.WaitGroup) {
	fmt.Println(barang, "Buah Sampah")
	wg.Done()
}
