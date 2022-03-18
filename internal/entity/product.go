package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name         string `gorm:"type:varchar(255);not null"`
	ExpenseRatio string `gorm:"type:varchar(255);not null"`

	Company uint `gorm:"type:varchar(255);not null"`

	Plan string `gorm:"type:text"`
	Rule string `gorm:"type:text"`

	Steps string `gorm:"type:text"`
}
