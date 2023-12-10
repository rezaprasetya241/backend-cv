package models

type User struct {
	Id          int64         `gorm:"primaryKey" json:"id"`
	Name        string        `gorm:"varchar(300)" json:"name"`
	Email       string        `gorm:"varchar(300)" json:"email"`
	Phone       int64         `gorm:"integer" json:"phone"`
	Experiences []Experience  `json:"experiences"`
	Certificate []Certificate `json:"certificate"`
	Education   []Education   `json:"education"`
	Skill       []Skill       `json:"skill"`
	Portofolio  []Portofolio  `json:"portofolio"`
}
