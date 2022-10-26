package models

import (
	"context"
	"errors"
	"time"

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

func GetMenuByID(uid int) (Menu, error) {
	var menu Menu
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := database.DB.WithContext(ctx).Model(Menu{}).Where("id = ?", uid).Take(&menu).Error; err != nil {
		return menu, errors.New("Menu not found!")
	}
	return menu, nil
}
