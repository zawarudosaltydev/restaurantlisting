package controllers

import (
	"encoding/json"
	"net/http"
)

// Restaurant Struct
type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Init restaurants var as a slice Restaurant struct
var restaurants []Restaurant

func Index(w http.ResponseWriter, r *http.Request) {
	// Mock Data - @todo - implement DB
	restaurants = append(restaurants, Restaurant{ID: "1", Name: "Swatow"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
