package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Claim is the structure to process the JWT
type Claim struct {
	Id             string `json:"id"`
	NombreCompleto string `json:"nombreCompleto"`
	Correo         string `json:"correo"`
	Rol            int    `json:"rol"`
	// StandardClaims -> expiration date of the token
	jwt.StandardClaims
}
