package models

import (
	"final/helpers"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	ID          int           `gorm:"primaryKey" json:"id"`
	Username    string        `gorm:"not null;uniqueIndex" form:"username" json:"username" valid:"required~Your username is required"`
	Email       string        `gorm:"not null;uniqueIndex" form:"email" json:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password    string        `gorm:"not null" form:"password" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age         int           `gorm:"not null" form:"age" json:"age" valid:"required~Your age is required"`
	CreatedAt   time.Time     `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt   time.Time     `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Photo"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Comment"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"SocialMedia"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age < 8 {
		return
	}

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age < 8 {
		return
	}
	if errCreate != nil {
		err = errCreate
		return
	}

	// u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
