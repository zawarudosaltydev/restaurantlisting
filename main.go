package main

import (
	"fmt"
	"github.com/zawarudosaltydev/restaurantlisting/controllers"
	"net/http"
)

func main() {
	// Route Handlers / Endpoints
	http.HandleFunc("/api/restaurants", controllers.Index)
	http.HandleFunc("/", controllers.Welcome)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
