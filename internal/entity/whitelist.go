package entity

import "gorm.io/gorm"

// Whitelist 白名单
type Whitelist struct {
	gorm.Model
	ProductID uint    `gorm:"primaryKey;autoIncrement:false"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`

	Name   string `gorm:"type:varchar(255);not null"`
	Number string `gorm:"type:varchar(255);not null"`
}
