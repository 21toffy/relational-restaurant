package models

import (
	"errors"
	"time"

	"github.com/21toffy/relational-restaurant/database"
	"github.com/21toffy/relational-restaurant/helpers"
	"golang.org/x/crypto/bcrypt"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	Id         uint      `json:"id" gorm:"primary_key"`
	Uid        string    `json:"uid" gorm:"unique"`
	Name       string    `json:"name"`
	Email      string    `json:"email" gorm:"size:255;not null;unique" validate:"required, unique"`
	Phone      string    `json:"phone" gorm:"default"`
	Address    string    `json:"address"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Password   string    `json:"password"`
}

type UserDisplay struct {
	Id      uint   `json:"id"`
	Uid     string `json:"uid"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}

func GetAllUsers(user *[]User) (err error) {
	if err = database.DB.Find(user).Error; err != nil {
		return err

	}
	return nil
}

func CreateUser(user *User) (err error) {
	result := database.DB.Where("email = ?", user.Email).Take(&user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		msg := "email already exists"
		ErrRegistered := errors.New(msg)
		return ErrRegistered
	}
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginCheck(email string, password string) (string, error) {
	var err error
	u := User{}
	err = database.DB.Model(User{}).Where("email = ?", email).Take(&u).Error
	if err != nil {
		return "", err
	}
	err = helpers.VerifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := helpers.GenerateToken(u.Id, u.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserByID(uid uint) (User, error) {
	var u User
	if err := database.DB.Model(User{}).Where("id = ?", uid).Take(&u).Error; err != nil {
		return u, errors.New("User not found!")
	}
	u.PrepareGive()
	return u, nil
}

func GetCurrentID(uid uint) (UserDisplay, error) {
	var u User
	var userDisplay UserDisplay

	if err := database.DB.Model(User{}).Where("id = ?", uid).Take(&u).Error; err != nil {
		return userDisplay, errors.New("User not found!")
	}
	userDisplay.Id = u.Id
	userDisplay.Uid = u.Uid
	userDisplay.Name = u.Name
	userDisplay.Email = u.Email
	userDisplay.Phone = u.Phone
	userDisplay.Address = u.Address
	return userDisplay, nil
}

func (u *User) PrepareGive() UserDisplay {
	var userDisplay UserDisplay

	userDisplay.Id = u.Id
	userDisplay.Uid = u.Uid
	userDisplay.Name = u.Name
	userDisplay.Email = u.Email
	userDisplay.Phone = u.Phone
	userDisplay.Address = u.Address
	return userDisplay
}

// func GetCurrentID(uid uint) (User, error) {
// 	var u User
// 	if err := database.DB.Model(User{}).Where("id = ?", uid).Take(&u).Error; err != nil {
// 		return u, errors.New("User not found!")
// 	}
// 	u.PrepareGive()
// 	return u, nil
// }

// func (u *User) PrepareGive() {
// 	u.Password = ""
// }

// //create a user
// func CreateUser(db *gorm.DB, User *User) (err error) {
// 	err = db.Create(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //get users
// func GetUsers(db *gorm.DB, User *[]User) (err error) {
// 	err = db.Find(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //get user by id
// func GetUser(db *gorm.DB, User *User, id int) (err error) {
// 	err = db.Where("id = ?", id).First(User).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// //update user
// func UpdateUser(db *gorm.DB, User *User) (err error) {
// 	db.Save(User)
// 	return nil
// }

// //delete user
// func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
// 	db.Where("id = ?", id).Delete(User)
// 	return nil
// }
