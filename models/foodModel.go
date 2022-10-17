package models

import (
	"time"

	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type Food struct {
	Id          uint        `json:"id" gorm:"primary_key"`
	Name        *string     `json:"name" validate:"required,min=3, max=50"`
	Price       *float64    `json:"price" validation:"required"`
	Description *string     `json:"food_description" gorm:"size:255;not null"`
	Food_image  *string     `json:"food_image" validate:"required"`
	Created_at  time.Time   `json:"created_at"`
	Updated_at  time.Time   `json:"updated_at"`
	Food_id     string      `json:"food_id" `
	UserId      int         `gorm:"default:null;"`
	User        UserDisplay `json:"user" gorm:"embedded;;embeddedPrefix:created_ ;foreignKey:UserId;association_foreignkey:ID"`
}

func (b *Food) TableName() string {
	return "food"
}

// func CreateFood(food *Food) (err error) {
// 	if err = database.DB.Create(food).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func CreateFood(food *Food) (err error) {
	if err = database.DB.Create(food).Error; err != nil {
		return err
	}
	return nil
}

func GetAllFoods(food *[]Food) (err error) {
	if err = database.DB.Find(food).Error; err != nil {
		return err

	}
	return nil
}
