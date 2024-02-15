package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/beevik/etree"
)

type M map[string]interface{}

type Food struct {
	Calories    string
	Description string
	Name        string
	Price       string
}

func xmlParser() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("simple.xml"); err != nil {
		log.Fatal(err.Error())
	}

	root := doc.SelectElement("breakfast_menu")
	rows := make([]M, 0)

	for _, valFood := range root.FindElements("food") {
		row := make(M)
		row["name"] = valFood.SelectElement("name").Text()
		row["price"] = valFood.SelectElement("price").Text()
		row["description"] = valFood.SelectElement("description").Text()
		row["calories"] = valFood.SelectElement("calories").Text()

		rows = append(rows, row)
	}

	bts, err := json.MarshalIndent(rows, "", "  ")

	if err != nil {
		log.Fatal(err)
	}

	foodStruct := []Food{}
	json.Unmarshal(bts, &foodStruct)

	for val := range foodStruct {
		dataInner := foodStruct[val]
		fmt.Println(val)
		fmt.Println("Calories : " + dataInner.Calories)
		fmt.Println("Description : " + dataInner.Description)
		fmt.Println("Name : " + dataInner.Name)
		fmt.Println("Price : " + dataInner.Price)
	}

}
