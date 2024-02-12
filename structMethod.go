package main

type Employee struct {
	EmployeeID int
	Nama       string
	Address    string
	Salary     float64
}

type Department struct {
	DepartmentID   string
	DepartmentName string
}

type Operasi interface {
	getAll() interface{}
	getOne() interface{}
}

func (emp Employee) getAll() []Employee {
	adi := Employee{
		EmployeeID: 123,
		Nama:       "Adi",
		Address:    "Jakarta",
		Salary:     100.55,
	}
	budi := Employee{
		EmployeeID: 456,
		Nama:       "Budi",
		Address:    "Bogor",
		Salary:     70.79,
	}
	hasil := []Employee{}
	hasil = append(hasil, adi, budi)
	return hasil
}

func (dept Department) getAll() []Department {
	finance := Department{
		DepartmentID:   "1",
		DepartmentName: "Finance",
	}
	hr := Department{
		DepartmentID:   "2",
		DepartmentName: "HR",
	}
	hasil := []Department{}
	hasil = append(hasil, finance, hr)
	return hasil
}
