package main

func main() {
	// Interface
	// emp := Employee{}
	// resultEmp := Employee.getAll(emp)

	// for _, val := range resultEmp {
	// 	fmt.Printf("%d | %s | %s | %f", val.EmployeeID, val.Nama, val.Address, val.Salary)
	// 	fmt.Println()
	// }

	// dept := Department{}
	// resultDept := Department.getAll(dept)

	// for _, val := range resultDept {
	// 	fmt.Printf("%s | %s", val.DepartmentID, val.DepartmentName)
	// 	fmt.Println()
	// }

	// emp := karyawan{}
	// getAll(&emp)
	// dept := departemen{}
	// getAll(&dept)

	// String.go
	// substring("Hello World", 0, 4)
	// splitString("Hello World Dunia")
	// stringReplace("HelloWorld", "o", "a", 2)
	// strPadLeft("agoes")

	// Encode b64 and SHA1
	// var text = "agoes"
	// fmt.Println(ValidMAC(text, "flp28"))
	// fmt.Println(tripleDES(text))
	// encodeDecodeBase64(text)
	// fmt.Printf("original : %s\n\n", text)

	// var hashed1, salt1 = doHashUsingSalt(text)
	// fmt.Printf("hashed 1 : %s\n\n", hashed1)

	// var hashed2, salt2 = doHashUsingSalt(text)
	// fmt.Printf("hashed 2 : %s\n\n", hashed2)

	// var hashed3, salt3 = doHashUsingSalt(text)
	// fmt.Printf("hashed 3 : %s\n\n", hashed3)

	// _, _, _ = salt1, salt2, salt3

	// Basic Channel
	// implGoroutine()
	// channelTest()
	// channelTestAgain()
	// waitGroup()

	// Non Goroutine
	// nonGoroutine()

	// Goroutine Channel
	// channelRoutine()

	// Goroutine WG
	// wgRoutine()

	// JSON & XML Parser
	// jsonParser()
	// xmlParser()
	token := getToken()
	getData(token)
}
