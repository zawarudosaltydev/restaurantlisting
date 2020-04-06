package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/zawarudosaltydev/restaurantlisting/models"
	"github.com/zawarudosaltydev/restaurantlisting/utils"
)

// IndexOrCreate will handle get all restaurants and create a new restaurant
func IndexOrCreate(w http.ResponseWriter, r *http.Request) {
	resp := utils.NewRespMsg()

	switch r.Method {
	case http.MethodGet:
		getRestaurants(resp, w)
	case http.MethodPost:
		createRestaurant(resp, w, r)
	default:
		resp.Code = http.StatusMethodNotAllowed
		w.Write(resp.JSONBytes())
	}
}

func getRestaurants(resp *utils.RespMsg, w http.ResponseWriter) {
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

func getRestaurant(id string, resp *utils.RespMsg, w http.ResponseWriter) {
	restaurant, err := models.GetRestaurant(id)
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

func updateRestaurant(id string, resp *utils.RespMsg, w http.ResponseWriter, r *http.Request) {
	var body map[string]*string
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = "server error"
		w.Write(resp.JSONBytes())
		return
	}
	err = models.UpdateRestaurant(id, body)
	if err != nil {
		fmt.Println(err)
		resp.Code = http.StatusInternalServerError
		resp.Message = "server error"
		w.Write(resp.JSONBytes())
		return
	}
	resp.Code = http.StatusOK
	w.Write(resp.JSONBytes())
	return
}

func createRestaurant(resp *utils.RespMsg, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var restaurant models.Restaurant
	err := json.NewDecoder(r.Body).Decode(&restaurant)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = "create a restaurant failed"
		w.Write(resp.JSONBytes())
		return
	}

	err = models.CreateRestaurant(restaurant)
	if err != nil {
		resp.Code = http.StatusInternalServerError
		resp.Message = "create a restaurant failed"
		w.Write(resp.JSONBytes())
		return
	}
	resp.Code = http.StatusOK
	resp.Message = "create a restaurant succeeded"
	w.Write(resp.JSONBytes())
}

// Restaurant will handle get/update/delete one restaurant
func Restaurant(w http.ResponseWriter, r *http.Request) {
	resp := utils.NewRespMsg()
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	switch r.Method {
	case http.MethodGet:
		getRestaurant(id, resp, w)
	case http.MethodPut:
		updateRestaurant(id, resp, w, r)
	case http.MethodDelete:
	default:
		resp.Code = http.StatusMethodNotAllowed
		w.Write(resp.JSONBytes())
	}

}
