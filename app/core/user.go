package core

import (
	"backend-go/app/models"
	"database/sql"
)

var users = []models.User{
	{
		Username: "user1",
		FullName: "User One",
		Email:    "user1@example.com",
		IsActive: true,
	},
	{
		Username: "user2",
		FullName: "User Two",
		Email:    "user2@example.com",
	},
}

func ListUsers() []models.User {
	return users
}

func GetUser(username string) (*models.User, error) {
	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, sql.ErrNoRows
}

func CreateUser(user models.User) error {
	users = append(users, user)
	return nil
}
