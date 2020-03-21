package main

import (
	"fmt"
	"net/http"

	"github.com/zawarudosaltydev/restaurantlisting/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Welcome)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err.Error())
		panic(err)
	}
}
