package models

import (
	"errors"
	"fmt"

	"shestapalau.by/rest/db"
	"shestapalau.by/rest/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	fmt.Println(userId)
	fmt.Println(user.ID)

	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, user.Email)

	var retreivedPassword string
	err := row.Scan(&user.ID, &retreivedPassword)

	if err != nil {
		return err
	}

	passwordValid := utils.CheckPasswordHash(user.Password, retreivedPassword)

	if !passwordValid {
		return errors.New("Invalid credentials")
	}

	return nil
}
