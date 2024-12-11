package models

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	ID     uint `json:"id"`
	jwt.RegisteredClaims
}

type JwtRefreshClaims struct {
	ID     uint `json:"id"`
	jwt.RegisteredClaims
}
