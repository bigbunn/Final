package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	User      *User     //`gorm:"foreignKey"`
	Comment   []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Comment"`
}

type PhotoGetResponse struct {
	GormModel
	ID        int       `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url"`
	UserID    int       `gorm:"foreignKey" json:"user_id"`
	CreatedAt time.Time `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	User      struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	} `json:"User"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
