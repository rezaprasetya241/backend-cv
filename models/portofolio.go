package models

import "time"

type Portofolio struct {
	Id        int64      `gorm:"id" json:"id"`
	Name      string     `gorm:"name" json:"name"`
	ImageUrl  []ImageUrl `json:"image_url"`
	Link      string     `gorm:"link" json:"link"`
	UserId    int64      `form:"user_id" json:"user_id"`
	StartDate time.Time  `gorm:"startDate" json:"start_date"`
	EndDate   time.Time  `gorm:"endDate" json:"end_date"`
}
