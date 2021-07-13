package repository

import (
	"errors"

	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/model"
)

var categories []*model.Category

func CreateCategory(category *model.Category) int64 {
	database.DB.Create(&category)
	return category.Id
}

func GetAllCategories() []*model.Category {
	database.DB.Find(&categories)
	return categories
}

func GetCategoryById(Id int) (*model.Category, error) {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		return category, nil
	} else {
		return nil, errors.New("category not found")
	}
}

func DeleteCategoryById(Id int64) error {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Delete(&category)
	if category != nil {
		return nil
	} else {
		return errors.New("category not found")
	}
}

func UpdateCategoryById(Id int64, categoryUpdate *model.Category) error {
	category := new(model.Category)
	database.DB.Where("id = ?", Id).Find(&category)
	if category != nil {
		category = categoryUpdate
		database.DB.Save(&category)
		return nil
	} else {
		return errors.New("category not found")
	}
}
