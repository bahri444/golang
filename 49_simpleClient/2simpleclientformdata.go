package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type pengguna struct {
	Id   string
	Name string
	Umur int
}

var link = "http://localhost:8080"

func form_data(Id string) (pengguna, error) {
	var err error
	var client = &http.Client{}
	var data pengguna

	var param = url.Values{}
	param.Set("id", Id)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", link+"/user", payload)
	if err != nil {
		return data, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
func main() {
	var user1, err = form_data("B002")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("id :%s\tNama :%s\tUmur :%d\n", user1.Id, user1.Name, user1.Umur)
}
