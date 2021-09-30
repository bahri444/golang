package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type pengguna struct {
	Id   string
	Name string
	Umur int
}

var baseUrl = "http://localhost:8080"

func fetchUser() ([]pengguna, error) {
	var err error
	var client = &http.Client{}
	var data []pengguna

	request, err := http.NewRequest("POST", baseUrl+"/users", nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func main() {
	var users, err = fetchUser()
	if err != nil {
		fmt.Println("error !", err.Error())
		return
	}
	for _, each := range users {
		fmt.Printf("Id :%s\tName :%s\tUmur :%d\t\n", each.Id, each.Name, each.Umur)
	}
}
