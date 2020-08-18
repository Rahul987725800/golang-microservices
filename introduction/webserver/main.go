package main

import (
	"fmt"
	"net/http"
 	"github.com/Pallinder/go-randomdata"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([] byte(randomdata.SillyName()))
	})
	fmt.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}