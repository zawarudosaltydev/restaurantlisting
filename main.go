package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/zawarudosaltydev/restaurantlisting/controllers"
	"log"
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
	// Init Router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	restaurants = append(restaurants, Restaurant{ID: "1", Name: "Swatow"})

	// Route Handlers / Endpoints
	r.HandleFunc("/api/restaurants", getRestaurants).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

	http.HandleFunc("/", controllers.Welcome)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
