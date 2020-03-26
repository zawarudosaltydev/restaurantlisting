package models

import "fmt"

// Restaurant data struct
type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AllRestaurants return all the restaurant from database
func AllRestaurants() ([]Restaurant, error) {
	rows, err := db.Query("SELECT id, name FROM restaurants")
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	defer rows.Close()

	rts := make([]Restaurant, 0)
	for rows.Next() {
		rt := Restaurant{}
		err := rows.Scan(&rt.ID, &rt.Name)
		if err != nil {
			fmt.Printf(err.Error())
			return nil, err
		}
		rts = append(rts, rt)
	}

	if err = rows.Err(); err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return rts, nil
}
