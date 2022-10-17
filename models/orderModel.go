package models

import (
	"time"

	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type Order struct {
	Base
	Order_Date time.Time `json:"order_date" validate:"required"`
	Order_id   string    `json:"order_id"`
	Table_id   *string   `json:"table_id" validate:"required"`
}

func CreateOrder(order *Order) (err error) {
	if err = database.DB.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrders(order *[]Order) (err error) {
	if err = database.DB.Find(order).Error; err != nil {
		return err

	}
	return nil
}
