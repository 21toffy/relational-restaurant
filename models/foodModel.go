package models

import (
	"context"
	"errors"
	"time"

	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type Food struct {
	Id          int         `json:"id" gorm:"primary_key"`
	Name        *string     `json:"name" validate:"required,min=3, max=50"`
	Price       *float64    `json:"price" validation:"required"`
	Description *string     `json:"food_description" gorm:"size:255;not null"`
	Food_image  *string     `json:"food_image" validate:"required"`
	Created_at  time.Time   `json:"created_at"`
	Updated_at  time.Time   `json:"updated_at"`
	Food_id     string      `json:"food_id" `
	UserId      int         `gorm:"default:null;"`
	User        UserDisplay `json:"user" gorm:"embedded;embeddedPrefix:created_ ;foreignKey:UserId;association_foreignkey:ID"`
	Deleted     bool        `gorm:"default:false;"`
}

func (b *Food) TableName() string {
	return "food"
}

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

func GetFoodByID(uid int) (Food, error) {
	var food Food
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	if err := database.DB.WithContext(ctx).Model(Food{}).Where("id = ? AND deleted = ?", uid, false).Take(&food).Error; err != nil {
		return food, errors.New("Food not found!")
	}
	return food, nil
}
