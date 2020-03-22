package controllers

import (
	"encoding/json"
	"github.com/zawarudosaltydev/restaurantlisting/models"
	"net/http"
)

// Init restaurants var as a slice Restaurant struct
var restaurants []models.Restaurant

func Index(w http.ResponseWriter, r *http.Request) {
	// Mock Data - @todo - implement DB
	restaurants = append(restaurants, models.Restaurant{ID: "1", Name: "Swatow"})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
