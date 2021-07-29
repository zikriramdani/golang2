package entity

import "github.com/dgrijalva/jwt-go"

// User object for REST(CRUD)
type User struct {
	jwt.StandardClaims
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Group    string `json:"Group"`
}
