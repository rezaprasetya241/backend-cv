package models

type Skill struct {
	Id          int64  `gorm:"id" json:"id"`
	Name        string `gorm:"name" json:"name"`
	Description string `gorm:"description" json:"description"`
	UserId      int64  `form:"user_id" json:"user_id"`
}
