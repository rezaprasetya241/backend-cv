package models

type JobDetails struct {
	Id           int64  `gorm:"id" json:"id"`
	Name         string `gorm:"name" json:"name"`
	Description  string `gorm:"description" json:"description"`
	ExperienceId int64  `json:"experience_id" form:"experience_id"`
}
