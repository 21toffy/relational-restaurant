package models

import (
	"errors"
	"fmt"
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

func GetOrderByID(uid int) (Order, error) {
	var order Order
	if err := database.DB.Model(Order{}).Where("id = ?", uid).Take(&order).Error; err != nil {
		fmt.Print("000000000000000000000000000____", err, uid, "____000000000000000000000000")
		return order, errors.New("order not found")
	}
	return order, nil
}
