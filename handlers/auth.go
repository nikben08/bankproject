package handlers

import (
	"bankproject/config"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/hex"

	"github.com/golang-jwt/jwt"
)

func (h handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type CreateUserRequest struct {
		UserName    string `json:"username"`
		Password    string `json:"password"`
		AccessLevel int    `json:"access_level"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request CreateUserRequest
	json.Unmarshal(body, &request)

	user := User{Username: request.UserName,
		Hash:        generateHash(request.Password),
		AccessLevel: request.AccessLevel}

	if result := h.DB.Create(&user); result.Error != nil {
		json.NewEncoder(w).Encode("Couldnt create user")
	}

	token, _ := generateJWT(user.Username, user.ID, user.AccessLevel)

	type Response struct {
		Code    string
		Message string
		Token   string
	}

	response := Response{Code: "200", Message: "User successfully created", Token: token}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h handler) Login(w http.ResponseWriter, r *http.Request) {
	type LoginRequest struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var request LoginRequest
	json.Unmarshal(body, &request)

	user := User{Username: request.UserName, Hash: generateHash(request.Password)}
	found := User{}
	if result := h.DB.Where("username = ?", user.Username).First(&found); result.Error != nil {
		json.NewEncoder(w).Encode("Error")
	}
	if found.Hash == user.Hash {
		token, _ := generateJWT(found.Username, found.ID, found.AccessLevel)

		type Response struct {
			Code    string
			Message string
			Token   string
		}

		response := Response{Code: "200", Message: "Successfolly logged", Token: token}

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode("Wrong username or password")
	}
}

func generateJWT(username string, userId int, accessLevel int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":      userId,
		"username":    username,
		"accessLevel": accessLevel,
	})
	secretKey := []byte(config.Config("JWT_SECRET_FOR_LOCAL"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, err
}

func generateHash(pass string) string {
	h := sha256.New()
	h.Write([]byte(pass))
	bs := hex.EncodeToString(h.Sum(nil))
	return bs
}
