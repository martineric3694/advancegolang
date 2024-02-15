package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type data struct {
	Isi     []dataAll `json:"data"`
	Message string    `json:"message"`
	Nama    string    `json:"nama"`
	RCode   string    `json:"rCode"`
}

type postIn struct {
	No_va string
	Nama  string
	Reqid string
}
type dataAll struct {
	No_va string `json:"no_va"`
	Nama  string `json:"nama"`
}

// JSON Parser from URL
func jsonParser() {
	httpposturl := "http://192.168.200.194:1212/vape"
	// akses ke employee -> menambah header authorization Bearer -> getToken dulu kemudian baru tambahkan header
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData []byte

	jsonData = []byte(`{
		"no_va" : "212351",
    	"nama" : "agoesps",
    	"reqid" : "02"
	}`)
	postIn := postIn{}
	json.Unmarshal(jsonData, &postIn)
	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if postIn.Reqid == "03" {
		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			fmt.Println("Time Out")
			return
		}
		defer response.Body.Close()

		fmt.Println("response Status:", response.Status)
		if response.Status == "200 OK" {
			fmt.Println("response Headers:", response.Header)
			body, _ := io.ReadAll(response.Body)
			fmt.Println("response Body:", string(body))
			dataResp := data{}
			err := json.Unmarshal(body, &dataResp)
			if err != nil {
				fmt.Print(err)
				return
			}
			fmt.Println(dataResp.Message)
			fmt.Println(dataResp.Nama)
			fmt.Println(dataResp.RCode)
			return
		} else {
			// selain status 200
		}
	} else {
		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			fmt.Println("Time Out")
			return
		}
		defer response.Body.Close()

		fmt.Println("response Status:", response.Status)
		if response.Status == "200 OK" {
			fmt.Println("response Headers:", response.Header)
			body, _ := io.ReadAll(response.Body)
			dataResp := data{}
			err := json.Unmarshal(body, &dataResp)
			if err != nil {
				fmt.Print(err)
				return
			}

			dataIn := dataResp.Isi
			// i := 0
			for val := range dataIn {
				// i++
				dataInner := dataResp.Isi[val]
				fmt.Print(val)
				fmt.Println("." + dataInner.Nama + "(" + dataInner.No_va + ")")
			}

		} else {
			// selain status 200
		}
	}
}

// get data from API using generated token and bearer
func getData(token string) {
	httpposturl := "http://localhost:9001/employees"
	fmt.Println("HTTP JSON POST URL:", httpposturl)
	request, _ := http.NewRequest("POST", httpposturl, nil)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Add("Authorization", "Bearer "+token)
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
		return
	}
	defer response.Body.Close()

	fmt.Println("response Headers:", response.Header)
	fmt.Println("Authorization Headers:", "Bearer :"+token)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}

// get Token from API
func getToken() string {
	httpposturl := "http://localhost:9001/user/login"
	// akses ke employee -> menambah header authorization Bearer -> getToken dulu kemudian baru tambahkan header
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData []byte = []byte(`{
		"email":"martin@gmail.com",
		"password":"123456"
	}`)
	postIn := postIn{}
	json.Unmarshal(jsonData, &postIn)
	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
		return "failed"
	}
	defer response.Body.Close()

	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	var data map[string]string
	json.Unmarshal(body, &data)
	return data["token"]
}
