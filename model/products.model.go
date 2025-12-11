package model

type Products struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"type:varchar(200);not null" json:"name"`
	Price       float64 `gorm:"type:decimal;not null" json:"price"`
	Ratings     float32 `gorm:"type:float;not null" json:"ratings"`
	Description string  `gorm:"type:text;not null" json:"description"`
	Category    string  `gorm:"type:varchar(200);not null" json:"category"`
}
