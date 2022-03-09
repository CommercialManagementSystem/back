package entity

import "gorm.io/gorm"

type ProductAppendix struct {
	gorm.Model
	ProductID uint    `gorm:"type:unsigned int"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Url       string  `gorm:"type:varchar(255);not null"`
}
