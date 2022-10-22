package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	ID             int       `gorm:"primaryKey" json:"id"`
	Name           string    `gorm:"not null" json:"name" valid:"required~Your name is required"`
	SocialMediaUrl string    `gorm:"not null" json:"social_media_url" valid:"required~Your URL is required"`
	UserID         int       `json:"user_id"`
	UpdatedAt      time.Time `gorm:"not null;type:timestamp" json:"updated_at,omitempty"`
	CreatedAt      time.Time `gorm:"not null;type:timestamp" json:"created_at,omitempty"`
	User           *User
}

type SocialMediaGetResponse struct {
	SocialMedia struct {
		ID             int       `json:"id"`
		Name           string    `json:"name"`
		SocialMediaUrl string    `json:"social_media_url"`
		UserID         int       `json:"userId"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		User           struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
		} `json:"User"`
	} `json:"social_medias"`
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (s *SocialMedia) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
