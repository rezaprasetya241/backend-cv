package models

import "time"

type Education struct {
	Id         int64     `gorm:"id" json:"id"`
	StartDate  time.Time `gorm:"startDate" json:"start_date"`
	EndDate    time.Time `gorm:"endDate" json:"end_date"`
	SchoolName string    `gorm:"schoolName" json:"school_name"`
	Major      string    `gorm:"major" json:"major"`
	UserId     int64     `json:"user_id" form:"user_id"`
}
