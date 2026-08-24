package models

import "github.com/ndcuongg/event-booking-api.git/db"

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Save() error {
	query := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = id
	return err
}
