package models

type User struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Username    string `json:"username"`
	Hash        string `json:"hash"`
	AccessLevel int    `json:"access_level"`
}
