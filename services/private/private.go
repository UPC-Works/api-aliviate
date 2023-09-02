package private

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	models "github.com/UPC-Works/api-aliviate/models"
)

func CompareToken(password_found string, input_password string) error {

	//brypt works with slices of bytes (this password is not encrypted)
	passwordBytes := []byte(input_password)

	//Password is encrypted
	passwordBD := []byte(password_found)

	//Compare the password found in the Storage and the input password
	error_compare_hash := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if error_compare_hash != nil {
		return error_compare_hash
	}

	return nil
}

func DecryptJWT(input_jwt string, claims *models.Claim) (*jwt.Token, error) {

	tokenKey := []byte("TokenGenerasdsa$$asdas..23c1qweadorRestoner")

	token, error_parse := jwt.ParseWithClaims(input_jwt, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})
	if error_parse != nil {
		return token, nil
	}

	return token, nil
}

func EncryptPassword(input_password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(input_password), cost)
	return string(bytes), err
}

func GenerateJWT(input_id string, input_nombre string, input_apellido string, input_correo string) (string, error) {
	tokenKey := []byte("TokenGenerasdsa$$asdas..23c1qweadorRestoner")

	payload := jwt.MapClaims{
		"id":       input_id,
		"Nombre":   input_nombre,
		"Apellido": input_apellido,
		"Correo":   input_correo,
		"exp":      time.Now().Add(time.Hour * 1460).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//Se añade el string de firma para completar los 3 campos que se pide en http...
	tokenStr, err_signedString := token.SignedString(tokenKey)
	if err_signedString != nil {
		return tokenStr, err_signedString
	}

	return tokenStr, nil
}
