package models

import (
	"errors"

	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type OrderItem struct {
	Base
	Order_Item_id string  `json:"order_item_id" `
	Size          string  `json:"size" validate:"required,eq=S|eq=M|eq=L"`
	Quantity      int64   `json:"quantity" validate:"required,min=1"`
	Unit_price    float64 `json:"unit_price" validate:"required"`

	FoodId  int `json:"food_id" gorm:"default:null;"`
	OrderId int `json:"order_id" gorm:"default:null;"`

	Food  Food  `json:"food" gorm:"embedded;;embeddedPrefix:order_item_food_ ;foreignKey:Food;association_foreignkey:ID" `
	Order Order `json:"order" gorm:"embedded;;embeddedPrefix:order_item_order_ ;foreignKey:Order;association_foreignkey:ID"`
}

func CreateOrderItem(order_item *OrderItem) (err error) {
	if err = database.DB.Create(order_item).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOrderItem(order_item *[]OrderItem) (err error) {
	if err = database.DB.Find(order_item).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderItemsByOrder(order_uid string, order_item *[]OrderItem) error {

	m := make(map[string]interface{})
	m["order_id"] = order_uid
	if err := database.DB.Where(m).Find(order_item).Error; err != nil {
		return errors.New("Order not found!")
	}
	return nil
}
