package core

import (
	"errors"

	"github.com/danishjsheikh/go-backend-demo/app/models"
)

var (
	users = make(map[string]models.User)
)

// ListUsers returns a list of all users
func ListUsers() ([]models.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	var userList []models.User
	for _, user := range users {
		userList = append(userList, user)
	}
	if len(userList) == 0 {
		return nil, errors.New("No User not found")
	}
	return userList, nil
}

// GetUser retrieves a user by username
func GetUser(username string) (*models.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	user, exists := users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// CreateUser creates a new user
func CreateUser(user models.User) error {
	mutex.Lock()
	defer mutex.Unlock()

	users[user.Username] = user
	return nil
}

// UpdateUser updates an existing user
func UpdateUser(username string, updatedUser models.User) (*models.User, error) {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := users[username]
	if !exists {
		return nil, errors.New("user not found")
	}
	delete(users, username)
	users[updatedUser.Username] = updatedUser
	return &updatedUser, nil
}

// DeleteUser removes a user
func DeleteUser(username string) error {
	mutex.Lock()
	defer mutex.Unlock()

	_, exists := users[username]
	if !exists {
		return errors.New("user not found")
	}
	delete(users, username)
	return nil
}
