package models

import (
	"time"

	"github.com/ndcuongg/event-booking-api.git/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location"  binding:"required"`
	DateTime    time.Time `json:"date_time"  binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, date_time, user_id)
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var events []Event
	defer rows.Close()
	for rows.Next() {
		var e Event
		err = rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, err
}
