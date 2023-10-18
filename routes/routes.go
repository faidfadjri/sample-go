// HTTP

package routes

import (
	"example/go-api/controllers"
	"net/http"
)

func SetupRoutes() {
    http.HandleFunc("/users", controllers.GetUsers)
}