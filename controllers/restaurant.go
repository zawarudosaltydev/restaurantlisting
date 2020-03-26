package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zawarudosaltydev/restaurantlisting/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restaurants, err := models.AllRestaurants()

	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(restaurants)
}
