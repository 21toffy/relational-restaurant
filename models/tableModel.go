package models

import (
	"github.com/21toffy/relational-restaurant/database"
	_ "gorm.io/driver/postgres"
)

type Table struct {
	Base
	Table_id         string `json:"table_id" `
	Number_of_guests *int64 `json:"number_of_guests" validate:"required,min=2, max=50"`
	Table_number     *int64 `json:"table_number" validate:"required"`
}

func CreateTable(table *Table) (err error) {
	if err = database.DB.Create(table).Error; err != nil {
		return err
	}
	return nil
}

func GetAllTables(table *[]Table) (err error) {
	if err = database.DB.Find(table).Error; err != nil {
		return err

	}
	return nil
}
