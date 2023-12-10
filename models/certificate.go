package models

import "time"

type Certificate struct {
	Id          int64     `gorm:"id" json:"id"`
	Name        string    `gorm:"name" json:"name"`
	Description string    `gorm:"description" json:"description"`
	Date        time.Time `gorm:"date" json:"date"`
	UserId      int64     `form:"user_id" json:"user_id"`
}
