package main

import "fmt"

func main() {
	var data = []string{"saepul", "bahri", "bangemet", "epul"}
	for _, each := range data {
		func() {

			//recover untuk iife dalam bentuk perulangan
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("error occured", each, "| pesan :", r)
				} else {
					fmt.Println("application running is perfect")
				}
			}()
			panic("someone error")
		}()
	}
}
