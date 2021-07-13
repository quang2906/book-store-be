package repository

import (
	"errors"

	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/model"
)

var users []*model.User

func GetAllUsers() []*model.User {
	database.DB.Find(&users)
	return users
}

func GetUserById(Id int) (*model.User, error) {
	user := new(model.User)
	database.DB.Where("id = ?", Id).Find(&user)
	if user != nil {
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}
