package models

import (
	"gorm.io/gorm"
)

type Interest struct {
	gorm.Model
	BankID     uint    `json:"bank_id"`
	ID         int     `json:"id" gorm:"primaryKey"`
	Interest   float32 `json:"interest"`
	TimeOption int     `json:"time_option"`
	CreditType int     `json:"credit_type"`
}
