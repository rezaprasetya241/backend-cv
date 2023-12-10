package models

import "time"

type Experience struct {
	Id          int64        `gorm:"primaryKey" json:"id"`
	UserId      int64        `form:"user_id" json:"user_id"`
	StartDate   time.Time    `json:"start_date"`
	EndDate     time.Time    `json:"end_date"`
	Role        string       `gorm:"varchar(300)" json:"role"`
	Company     string       `gorm:"varchar(300)" json:"company"`
	Description string       `gorm:"text" json:"description"`
	JobDetails  []JobDetails `json:"job_details"`
}
