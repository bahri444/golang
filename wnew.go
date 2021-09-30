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

func getelementbyId(Id string) (pengguna, error) {
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
	var getuser, err = getelementbyId("G005")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Id :%s\tNama :%s\tUmur :%d\n", getuser.Id, getuser.Name, getuser.Umur)
}
