package models

import (
	"errors"
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

func GetOrderByID(id int) (Order, error) {
	var order Order
	if err := database.DB.Model(Order{}).Where("id = ?", id).Take(&order).Error; err != nil {

		return order, errors.New("order not found")
	}
	return order, nil
}

func GetOrderByUID(uid string) (Order, error) {
	var order Order
	if err := database.DB.Model(Order{}).Where("order_id = ?", uid).Take(&order).Error; err != nil {

		return order, errors.New("order not found")
	}
	return order, nil
}
