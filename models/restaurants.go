package models

import (
	"fmt"
	"strings"
	"time"
)

// Restaurant data struct
type Restaurant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Number    string    `json:"number"` //json: cannot unmarshal string into Go struct field Restaurant.number of type sql.NullString
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

// GetOneRestaurant will return the Restaurant by id
func GetOneRestaurant(id string) (Restaurant, error) {
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

// UpdateOneRestaurant will access database to update the Restaurant by id
func UpdateOneRestaurant(id string, body map[string]string) error {
	var args []interface{}
	sqlString := "UPDATE restaurants SET"
	for k, v := range body {
		sqlString += " " + k + "=?,"
		args = append(args, v)
	}
	sqlString = strings.TrimSuffix(sqlString, ",") + " WHERE id=?"
	args = append(args, id)

	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
