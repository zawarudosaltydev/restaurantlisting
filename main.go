package main

import (
	"encoding/json"
	"fmt"
	"github.com/zawarudosaltydev/restaurantlisting/controllers"
	"net/http"
)

// Restaurant Struct
type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Init restaurants var as a slice Restaurant struct
var restaurants []Restaurant

func getRestaurants(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}

func main() {
	// Mock Data - @todo - implement DB
	restaurants = append(restaurants, Restaurant{ID: "1", Name: "Swatow"})

	// Route Handlers / Endpoints
	http.HandleFunc("/api/restaurants", getRestaurants)
	http.HandleFunc("/", controllers.Welcome)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
