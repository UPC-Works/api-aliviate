package medico

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	medico_repository "github.com/UPC-Works/api-aliviate/repositories/medico"
)

func Update(c echo.Context) error {

	//Inicilization
	var input_medico *models.Medico

	//Bind the model
	err := c.Bind(&input_medico)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_medico.Id == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Debe enviar el Id",
			},
			Data: ""})
	}

	//Find in the Storage
	medico_found, error_findbyemail := medico_repository.Pg_FindOne(input_medico.Correo)
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

	//Storage the Medico
	update_medico := models.UpdateMedico(input_medico.Id, input_medico.Nombre, input_medico.Apellido, input_medico.Colegiatura, medico_found.Correo, medico_found.Contrasenia, input_medico.Direccion, input_medico.Especialidad)
	error_create_medico := medico_repository.Pg_Create(update_medico)
	if error_create_medico != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_medico.Error(),
			},
			Data: ""})
	}

	//OK
	return c.JSON(200, &helpers.ResponseString{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: "OK"})
}
