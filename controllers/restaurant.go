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

func getOneRestaurant(id string, resp *utils.RespMsg, w http.ResponseWriter) {
	restaurant, err := models.GetOneRestaurant(id)
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

func updateOneRestaurant(id string, body map[string]string, resp *utils.RespMsg, w http.ResponseWriter) {
	err := models.UpdateOneRestaurant(id, body)
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

// Restaurant will handle get/update/delete one restaurant
func Restaurant(w http.ResponseWriter, r *http.Request) {
	resp := utils.NewRespMsg()
	path := strings.Split(r.URL.Path, "/")
	id := path[len(path)-1]

	switch r.Method {
	case http.MethodGet:
		getOneRestaurant(id, resp, w)
	case http.MethodPut:
		var body map[string]string
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println(err)
			resp.Code = http.StatusInternalServerError
			resp.Message = "server error"
			w.Write(resp.JSONBytes())
			return
		}
		updateOneRestaurant(id, body, resp, w)
	case http.MethodDelete:
	default:
		resp.Code = http.StatusMethodNotAllowed
		w.Write(resp.JSONBytes())
	}

}
