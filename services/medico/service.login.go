package medico

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	medico_repository "github.com/UPC-Works/api-aliviate/repositories/medico"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func Login(c echo.Context) error {

	//Inicilization
	var input_medico *models.Medico

	//Bind the model
	err := c.Bind(&input_medico)
	if err != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//Validation of the Business Rules
	if len(input_medico.Correo) < 1 || len(input_medico.Correo) > 50 || len(input_medico.Contrasenia) < 8 || len(input_medico.Contrasenia) > 12 {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "El Correo no puede exceder la longitud de 50 y la Contraseña debe estar entre 8-12 caracteres",
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//Find in the Storage
	medico_found, error_findbyemail := medico_repository.Pg_FindOne("", input_medico.Correo)
	if error_findbyemail != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_findbyemail.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}
	if medico_found.Correo == "" {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Medico no encontrado",
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//Check the password
	error_compareToken := private_services.CompareToken(medico_found.Contrasenia, input_medico.Contrasenia)
	if error_compareToken != nil {
		return c.JSON(400, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Contraseña incorrecta, detalle:" + error_compareToken.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//Generate the token
	jwtKey, error_generatingJWT := private_services.GenerateJWT(medico_found.Id, medico_found.Nombre, medico_found.Apellido, medico_found.Correo, 2)
	if error_generatingJWT != nil {
		return c.JSON(500, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error when trying to generate the token, detail:" + error_generatingJWT.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//Storage in the cache
	err_add_re := medico_repository.Re_SetId(medico_found.Id)
	if err_add_re != nil {
		return c.JSON(500, &helpers.ResponseJwt{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   err_add_re.Error(),
			},
			Data: helpers.JwtStructure{
				JWT:            "",
				NombreCompleto: "",
				Correo:         "",
				Rol:            0,
			}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseJwt{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: helpers.JwtStructure{
			JWT:            jwtKey,
			NombreCompleto: medico_found.Nombre,
			Correo:         medico_found.Correo,
			Rol:            2,
		}})

}
