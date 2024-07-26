package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/padiazg/wasm-app-test/backend/models"
)

var users []models.User

func InitDB() {
	users = []models.User{
		{ID: "1", Username: "admin", Name: "Admin User", Email: "admin@example.com", Password: "admin123", Active: true},
	}
}

func GetUsers() ([]models.User, error) {
	return users, nil
}

func CreateUser(user *models.User) error {
	user.ID = fmt.Sprintf("%d", len(users)+1)
	users = append(users, *user)
	return nil
}

func GetUser(id string) (*models.User, error) {
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("User not found")
}

func UpdateUser(user *models.User) error {
	for i, u := range users {
		if u.ID == user.ID {
			users[i] = *user
			return nil
		}
	}
	return errors.New("User not found")
}

func DeleteUser(id string) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func AuthenticateUser(username, password string) (*models.User, error) {
	log.Printf("AuthenticateUser username:%s password:%s", username, password)
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user, nil
		}
	}
	return nil, errors.New("Invalid credentials")
}
