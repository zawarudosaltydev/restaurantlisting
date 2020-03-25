package controllers

import (
	"encoding/json"
	"github.com/zawarudosaltydev/restaurantlisting/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restaurants, err := models.AllRestaurants()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
