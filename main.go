package main

import (
	"fmt"
	"net/http"

	"github.com/zawarudosaltydev/restaurantlisting/controllers"
	"github.com/zawarudosaltydev/restaurantlisting/models"
)

func main() {
	models.InitDB("admin:admin@tcp(localhost:3308)/restaurantlisting")

	// Route Handlers / Endpoints
	http.HandleFunc("/api/restaurants", controllers.Index)
	http.HandleFunc("/", controllers.Welcome)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
