package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	User      *User     //`gorm:"foreignKey"`
	Photo     *Photo    //`gorm:"foreignKey"`
}

type CommentGetResponse struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	PhotoID   int       `json:"photo_id"`
	Message   string    `gorm:"not null" json:"message"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	User      struct {
		ID       int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"User"`
	Photo struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserID   int    `json:"user_id"`
	} `json:"Photo"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (c *Comment) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
