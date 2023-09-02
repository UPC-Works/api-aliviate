package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

//Claim is the structure to process the JWT
type Claim struct {
	Id          string `json:"id"`
	Country     int    `json:"country"`
	Rol         int    `json:"rol"`
	Full_name   string `json:"fullName"`
	SessionCode int    `json:"sessioncode"`
	Source      int    `json:"source"`
	// StandardClaims -> expiration date of the token
	jwt.StandardClaims
}
