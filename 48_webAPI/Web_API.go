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

var data = []pengguna{
	pengguna{"S001", "Saepul", 18},
	pengguna{"B002", "Bahri", 19},
	pengguna{"E003", "Epol", 20},
	pengguna{"B004", "Bang_emet", 21},
	pengguna{"G005", "Gant", 22},
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
func user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		var id = r.FormValue("id")
		for _, each := range data {
			if each.Id == id {
				result, err := json.Marshal(each)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(result)
				return
			}
		}
	}
	http.Error(w, "", http.StatusBadRequest)
}
func main() {
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
