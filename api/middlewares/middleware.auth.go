package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		//Sending the request
		resquest_http, _ := http.NewRequest("GET", "https://business-register-authentication.api-restoner.link/v1/auth", nil)
		resquest_http.Header.Add("auth-token", c.Request().Header.Get("auth-token"))
		client := &http.Client{}
		response_http, _ := client.Do(resquest_http)

		//Decoding the response
		var response_auth *helpers.ResponseAuth
		error_decode := json.NewDecoder(response_http.Body).Decode(&response_auth)
		if error_decode != nil {
			return c.JSON(403, &helpers.ResponseString{
				Error: helpers.ErrorStructure{
					Code:   9459,
					Detail: "Error in the auth, details: " + error_decode.Error(),
				},
				Data: ""})
		}
		if response_auth.Data.Id == "" {
			return c.JSON(403, &helpers.ResponseString{
				Error: helpers.ErrorStructure{
					Code:   9459,
					Detail: "Error in the auth, details: This user does not exist",
				},
				Data: ""})
		}

		//Assigning the values
		c.Set("id", response_auth.Data.Id)
		c.Set("nombre", response_auth.Data.Nombre)
		c.Set("apellido", response_auth.Data.Apellido)
		c.Set("correo", response_auth.Data.Correo)

		//OK
		return next(c)
	}

}
