package model

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	MovieId string `gorm:"not null" json:"movie_id"`
	Body    string `gorm:"not null" json:"body"`
	UserIp  string `gorm:"not null" json:"user_ip"`
}
