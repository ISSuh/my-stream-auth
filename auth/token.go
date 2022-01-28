package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type Claim struct {
	Email string
	jwt.StandardClaims
}

type Token struct {
	AccessToken  string `json: "access_token`
	RefreshToken string `json: "refresh_token"`
}
