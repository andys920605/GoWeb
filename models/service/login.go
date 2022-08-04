package service

import "github.com/golang-jwt/jwt"

type LoginReq struct {
	Account  string `json:"account,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// custom claims
type Claims struct {
	Message string `json:"message,omitempty"`
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}
