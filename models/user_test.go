package models

import (
	"testing"

	"gorm.io/gorm"
)

func TestUser_BeforeCreate(t *testing.T) {

	type args struct {
		tx *gorm.DB
	}
	tests := []struct {
		name    string
		u       *User
		args    args
		wantErr bool
	}{
		{
			name: "Minimum Age",
			u: &User{
				Username: "bani",
				Email:    "bani@gmail.com",
				Password: "123456",
				Age:      7,
			},
			wantErr: true,
		},
		{
			name: "Password not enough",
			u: &User{
				Username: "bani",
				Email:    "bani@gmail.com",
				Password: "5",
				Age:      9,
			},
			wantErr: true,
		},
		{
			name: "Email not valid",
			u: &User{
				Username: "bani",
				Email:    "banigmail.com",
				Password: "654321",
				Age:      9,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.BeforeCreate(tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("User.BeforeCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
