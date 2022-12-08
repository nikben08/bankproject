package handlers

import (
	"bankproject/models"

	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

type User models.User
type Bank models.Bank
