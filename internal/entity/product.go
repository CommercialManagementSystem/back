package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(255);not null"`
	ExpenseRatio float64 `gorm:"type:float;not null"`

	CompanyID uint `gorm:"type:unsigned int"`
	Company   User `gorm:"foreignKey:CompanyID;references:ID"`

	Plan string `gorm:"type:longtext"`
	Rule string `gorm:"type:longtext"`

	Steps string `gorm:"type:longtext"`
}
