package main

import "fmt"

type karyawan struct {
	EmployeeID int
	Nama       string
	Address    string
	Salary     float64
}

type departemen struct {
	DepartmentID   string
	DepartmentName string
}

type Operation interface {
	getAll()
}

func getAll(op Operation) {
	op.getAll()
}

func (e *karyawan) getAll() {
	adi := karyawan{
		EmployeeID: 123,
		Nama:       "Adi",
		Address:    "Jakarta",
		Salary:     100.55,
	}
	budi := karyawan{
		EmployeeID: 456,
		Nama:       "Budi",
		Address:    "Bogor",
		Salary:     70.79,
	}
	fmt.Print(adi)
	fmt.Print("+")
	fmt.Print(budi)
	// hasil := []data{}
	// hasil = append(hasil, adi, budi)
	// return hasil
}

func (d *departemen) getAll() {
	finance := departemen{
		DepartmentID:   "1",
		DepartmentName: "Finance",
	}
	hr := departemen{
		DepartmentID:   "2",
		DepartmentName: "HR",
	}
	hasil := []departemen{}
	hasil = append(hasil, finance, hr)
	fmt.Println(hasil)
}
