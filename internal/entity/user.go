package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UID       string `gorm:"column:uid;type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Authority int    `gorm:"type:int;default:0;not null"`

	Name  string `gorm:"type:varchar(255);not null"`
	Sex   int    `gorm:"type:int;not null"`
	Phone string `gorm:"type:varchar(255);not null"`
	Email string `gorm:"type:varchar(255);not null"`

	RoleID uint `gorm:"not null"`
	Role   Role `gorm:"foreignKey:RoleID;references:ID"`
}
