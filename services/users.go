package services

import (
	"errors"

	"fanchann/elegant_err/models"
)

type UsersServices struct{}

func NewUsersServices() *UsersServices {
	return &UsersServices{}
}

func (u *UsersServices) SearchUserById(id int) (*models.UsersModels, error) {
	for _, user := range models.FakeData {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
