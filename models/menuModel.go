package models

import (
	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type Menu struct {
	Base
	Name     string `json:"name" validate:"required, min=3, max=100"`
	Category string `json:"category" validate:"required;eq=rood|eq=drink"`
	Menu_id  string `json:"food_id"`
}

func CreateMenu(menu *Menu) (err error) {
	if err = database.DB.Create(menu).Error; err != nil {
		return err
	}
	return nil
}

func GetAllMenus(menu *[]Menu) (err error) {
	if err = database.DB.Find(menu).Error; err != nil {
		return err

	}
	return nil
}
