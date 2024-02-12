package main

import "fmt"

// Case
// restoran -> checking belanjaan (ikan, daging, sayur, buah, bawang)
//
// ikan -> cuci -> simpan frozen
// daging -> potong -> simpan frozen
// sayur -> kupas -> simpan box -> simpan chiler
// buah -> cuci -> simpan chiler
// bawang -> kupas -> simpan box -> simpan chiler
// jika bukan barang diatas -> buang sampah

func nonGoroutine() {
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
			cuci(brg)
			i++
		case "daging":
			fmt.Println("Menemukan", brg)
			potong(brg)
			i++
		case "sayur":
			fmt.Println("Menemukan", brg)
			kupas(brg)
			i++
		case "buah":
			fmt.Println("Menemukan", brg)
			cuci(brg)
			i++
		case "bawang":
			fmt.Println("Menemukan", brg)
			kupas(brg)
			i++
		default:
			buangSampah(brg)
			i++
		}
	}
}

func cuci(barang string) {
	fmt.Println("Cuci", barang)
	if barang == "ikan" {
		simpanFrozen(barang)
	} else {
		simpanChiler(barang)
	}
}

func potong(barang string) {
	fmt.Println("Potong", barang)
	simpanFrozen(barang)
}

func kupas(barang string) {
	fmt.Println("Cuci", barang)
	simpanBox(barang)
	simpanChiler(barang)
}

func simpanBox(barang string) {
	fmt.Println(barang, "Simpan Box")
}

func simpanFrozen(barang string) {
	fmt.Println(barang, "Simpan Frozen")
}

func simpanChiler(barang string) {
	fmt.Println(barang, "Simpan Chiler")
}

func buangSampah(barang string) {
	fmt.Println(barang, "Buah Sampah")
}
