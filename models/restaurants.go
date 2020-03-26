package models

import (
	"database/sql"
	"fmt"
)

// Restaurant data struct
type Restaurant struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Address string         `json:"address"`
	Number  sql.NullString `json:"number"`
}

// AllRestaurants return all the restaurant from database
func AllRestaurants() ([]Restaurant, error) {
	rows, err := db.Query("SELECT id, name, address, number FROM restaurants")
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	defer rows.Close()

	rts := make([]Restaurant, 0)
	for rows.Next() {
		rt := Restaurant{}
		err := rows.Scan(
			&rt.ID,
			&rt.Name,
			&rt.Address,
			&rt.Number,
		)
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
