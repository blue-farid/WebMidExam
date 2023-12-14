package model

import (
	"gorm.io/gorm"
)

type Basket struct {
	gorm.Model
	Data   string `json:"data" gorm:"not null"`
	State  string `json:"state" gorm:"not null"`
	UserID uint
	User   User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (b *Basket) CheckOwnership(uID uint) bool {
	return uID == b.UserID
}
