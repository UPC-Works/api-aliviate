package admin

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	admin_repository "github.com/UPC-Works/api-aliviate/repositories/admin"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func Login(c echo.Context) error {

	//Inicilization
	var input_admin *models.Admin

	//Bind the model
	err := c.Bind(&input_admin)
	if err != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   9451,
				Detail: "Revisa la estructura o el tipo del valor",
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}

	//Validation of the Business Rules
	if len(input_admin.Correo) < 1 || len(input_admin.Correo) > 50 || len(input_admin.Contrasenia) < 8 || len(input_admin.Contrasenia) > 12 {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   9452,
				Detail: "El Correo no puede exceder la longitud de 50 y la Contraseña debe estar entre 8-12 caracteres",
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}

	//Find in the Storage
	admin_found, error_findbyemail := admin_repository.Pg_FindOne(input_admin.Correo)
	if error_findbyemail != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   9457,
				Detail: error_findbyemail.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}
	if admin_found.Correo == "" {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   9455,
				Detail: "Admin no encontrado",
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}

	//Check the password
	error_compareToken := private_services.CompareToken(admin_found.Contrasenia, admin_found.Contrasenia)
	if error_compareToken != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   9459,
				Detail: "Contraseña incorrecta, detalle:" + error_compareToken.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}

	//Generate the token
	jwtKey, error_generatingJWT := private_services.GenerateJWT(admin_found.Id, admin_found.NombreCompleto, "", admin_found.Correo, 1)
	if error_generatingJWT != nil {
		return c.JSON(500, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				Code:   94510,
				Detail: "Error when trying to generate the token, detail:" + error_generatingJWT.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:      "",
				Nombre:   "",
				Apellido: "",
				Correo:   "",
			}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseJwt{
		Error: helpers.ErrorStructure{
			Code:   0,
			Detail: "",
		},
		Data: helpers.JwtStructure{
			JWT:      jwtKey,
			Nombre:   admin_found.NombreCompleto,
			Apellido: "",
			Correo:   admin_found.Correo,
		}})

}
