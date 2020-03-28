package main

import (
	"fmt"
	"net/http"

	"github.com/zawarudosaltydev/restaurantlisting/controllers"
	"github.com/zawarudosaltydev/restaurantlisting/models"
)

const PORT = ":8080"

func main() {
	models.InitDB("admin:admin@tcp(localhost:3308)/restaurantlisting?parseTime=true")

	// Route Handlers / Endpoints
	http.HandleFunc("/api/restaurants", controllers.Index)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
