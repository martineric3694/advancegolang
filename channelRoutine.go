package main

import (
	"fmt"
	"runtime"
)

var cuciBarangChannel = make(chan string)
var potongBarangChannel = make(chan string)
var kupasBarangChannel = make(chan string)
var simpanBoxBarangChannel = make(chan string)
var simpanFrozenBarangChannel = make(chan string)
var simpanChilerBarangChannel = make(chan string)
var buangSampahBarangChannel = make(chan string)

func channelRoutine() {

	runtime.GOMAXPROCS(2)

	belanja := []string{
		"ikan", "daging", "plastik", "sayur", "plastik", "bumbu", "buah", "sterofoam", "bawang", "ikan",
		"ikan", "daging", "plastik", "sayur", "ikan", "daging", "plastik", "sayur", "bawang", "buah",
		"sterofoam", "bawang", "daging", "plastik", "bawang", "buah", "sayur", "bawang", "buah", "sterofoam", "ikan",
	}

	go cuciBarang()
	go potongBarang()
	go kupasBarang()
	go simpanBoxBarang()
	go simpanFrozenBarang()
	go simpanChilerBarang()
	go buangSampahBarang()

	i := 1
	for _, brg := range belanja {
		fmt.Print(i)
		fmt.Print(" ")
		switch brg {
		case "ikan":
			fmt.Println("Menemukan", brg)
			cuciBarangChannel <- brg
			i++
		case "daging":
			fmt.Println("Menemukan", brg)
			potongBarangChannel <- brg
			i++
		case "sayur":
			fmt.Println("Menemukan", brg)
			kupasBarangChannel <- brg
			i++
		case "buah":
			fmt.Println("Menemukan", brg)
			cuciBarangChannel <- brg
			i++
		case "bawang":
			fmt.Println("Menemukan", brg)
			kupasBarangChannel <- brg
			i++
		default:
			buangSampahBarangChannel <- brg
			i++
		}
	}
}

func cuciBarang() {
	for itemCuci := range cuciBarangChannel {
		fmt.Println("Cuci", itemCuci)
		if itemCuci == "ikan" {
			simpanFrozenBarangChannel <- itemCuci
		} else {
			simpanChilerBarangChannel <- itemCuci
		}
	}
}

func potongBarang() {
	for itemPotong := range potongBarangChannel {
		fmt.Println("Potong", itemPotong)
		simpanFrozenBarangChannel <- itemPotong
	}
}

func kupasBarang() {
	for itemKupas := range kupasBarangChannel {
		fmt.Println("Cuci", itemKupas)
		simpanBoxBarangChannel <- itemKupas
		simpanChilerBarangChannel <- itemKupas
	}
}

func simpanBoxBarang() {
	for itemSimpanBox := range simpanBoxBarangChannel {
		fmt.Println(itemSimpanBox, "Simpan Box")
	}
}

func simpanFrozenBarang() {
	for itemSimpanFrozen := range simpanFrozenBarangChannel {
		fmt.Println(itemSimpanFrozen, "Simpan Frozen")
	}
}

func simpanChilerBarang() {
	for itemSimpanChiler := range simpanChilerBarangChannel {
		fmt.Println(itemSimpanChiler, "Simpan Chiler")
	}
}

func buangSampahBarang() {
	for itemBuangSampah := range buangSampahBarangChannel {
		fmt.Println(itemBuangSampah, "Buah Sampah")
	}
}
