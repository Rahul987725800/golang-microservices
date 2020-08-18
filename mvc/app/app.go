package app

import (
	"net/http"
	"github.com/Rahul987725800/golang-microservices/mvc/controllers"
	"fmt"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
	fmt.Println("starting server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}