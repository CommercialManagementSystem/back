package entity

type ProductUser struct {
	ProductID uint    `gorm:"primaryKey;autoIncrement:false"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID"`

	UserID uint `gorm:"primaryKey;autoIncrement:false"`
	User   User `gorm:"foreignKey:UserID;references:ID"`
}
