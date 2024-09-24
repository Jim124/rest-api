package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "insert into users(email,password)values(?,?)"
	stmt, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer stmt.Close()
	hashPassword, error := utils.GenerateHashPassword(u.Password)
	if error != nil {
		return error
	}
	row, error := stmt.Exec(u.Email, hashPassword)
	if error != nil {
		return error
	}
	id, error := row.LastInsertId()
	u.ID = id
	return error
}

func (u *User) ValidateCredentials() error {
	query := "select id,password from users where email=?"
	row := db.DB.QueryRow(query, u.Email)
	var hashPassword string
	error := row.Scan(&u.ID, &hashPassword)
	if error != nil {
		return errors.New("credentials invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(hashPassword, u.Password)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
