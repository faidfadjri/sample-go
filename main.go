package main

import (
	"example/go-api/routes"
	"fmt"
	"net/http"
)

func main() {
	routes.SetupRoutes()
    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}