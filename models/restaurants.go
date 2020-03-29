package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Restaurant data struct
type Restaurant struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Number    sql.NullString `json:"number"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// AllRestaurants return all the restaurant from database
func AllRestaurants() ([]Restaurant, error) {
	rows, err := db.Query("SELECT * FROM restaurants")
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
			&rt.CreatedAt,
			&rt.UpdatedAt,
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

// OneRestaurant will return the Restaurant by id
func OneRestaurant(id string) (Restaurant, error) {
	rt := Restaurant{}
	stmt, err := db.Prepare("SELECT * FROM restaurants WHERE id = ? LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return rt, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(
		&rt.ID,
		&rt.Name,
		&rt.Address,
		&rt.Number,
		&rt.CreatedAt,
		&rt.UpdatedAt,
	)
	if err != nil {
		fmt.Println(err.Error())
		return rt, err
	}

	return rt, nil
}
