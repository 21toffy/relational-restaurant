package models

import (
	"time"

	_ "gorm.io/driver/postgres"
)

type Base struct {
	Id        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Deleted   bool       `gorm:"default:false;"`
}
