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
}

// BeforeCreate will set a UUID rather than numeric ID.
// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	fmt.Print("uuuuuuuuiiiiiiiiddddd", "uuid")
// 	uuid := helpers.GenerateUUID()
// 	fmt.Print(uuid, "uuid")
// 	return scope.SetColumn("ID", uuid)
// }
