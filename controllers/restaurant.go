package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/zawarudosaltydev/restaurantlisting/models"
	"github.com/zawarudosaltydev/restaurantlisting/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	resp := utils.NewRespMsg()
	restaurants, err := models.AllRestaurants()

	if err != nil {
		fmt.Println(err.Error())
		resp.Code = http.StatusInternalServerError
		resp.Message = "server error"
		w.Write(resp.JSONBytes())
		return
	}

	resp.Code = http.StatusOK
	resp.Data = restaurants
	w.Write(resp.JSONBytes())
}

// GetRestaurant return one restaurant
func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	resp := utils.NewRespMsg()
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	restaurant, err := models.OneRestaurant(id)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.Code = http.StatusNoContent
			resp.Message = "not existing"
		} else {
			resp.Code = http.StatusInternalServerError
			resp.Message = "server error"
		}
		w.Write(resp.JSONBytes())
		return
	}

	resp.Code = http.StatusOK
	resp.Data = restaurant
	w.Write(resp.JSONBytes())
}
