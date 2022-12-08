package handlers

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

func (h handler) ParseJwt(tokenString string) int {
	secretKey := []byte("0d1ea4c256cd50a2a7ccc7c65ee4e21df06d")
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		found := User{}
		if result := h.DB.Where("username = ?", claims["username"].(string)).First(&found); result.Error != nil {
			return 404
		}
		return found.ID
	} else {
		return 404
	}
}
