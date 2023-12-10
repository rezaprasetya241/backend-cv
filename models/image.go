package models

type ImageUrl struct {
	Id           int64  `gorm:"id" json:"id"`
	Url          string `gorm:"url" json:"url"`
	PortofolioId int64  `form:"portofolio_id" json:"portofolio_id"`
}
