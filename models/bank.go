package models

import (
	"gorm.io/gorm"
)

type Bank struct {
	gorm.Model
	ID       int        `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name"`
	Interest []Interest `gorm:"foreignKey:BankID; constraint:OnDelete:CASCADE;" json:"-"`
}
