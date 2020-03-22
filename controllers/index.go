package controllers

import (
	"encoding/json"
	"github.com/zawarudosaltydev/restaurantlisting/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Mock Data - @todo - implement DB
	restaurant := models.Restaurant{ID: "1", Name: "Swatow"}
	restaurants := []models.Restaurant{restaurant}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
