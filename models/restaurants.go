package models

type Restaurant struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func AllRestaurants() ([]*Restaurant, error) {
	rows, err := db.Query("SELECT * FROM restaurants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	rts := make([]*Restaurant, 0)
	for rows.Next() {
		rt := new(Restaurant)
		err := rows.Scan(&rt.ID, &rt.Name)
		if err != nil {
			return nil, err
		}
		rts = append(rts, rt)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rts, nil
}
