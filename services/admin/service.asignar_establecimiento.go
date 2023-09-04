package admin

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	establecimiento_medico_repository "github.com/UPC-Works/api-aliviate/repositories/establecimiento_medico"
)

func AsignarEstablecimiento(c echo.Context) error {

	//Inicilization
	var input_establecimiento_medico *models.EstablecimientoMedico

	//Bind the model
	err := c.Bind(&input_establecimiento_medico)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_establecimiento_medico.IdEstablecimiento == "" || input_establecimiento_medico.IdMedico == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "IdEstablecimiento y IdMedico deben ser enviados",
			},
			Data: ""})
	}

	//Storage the Establecimiento-Medico
	error_create_paciente := establecimiento_medico_repository.Pg_Create(input_establecimiento_medico.IdMedico, input_establecimiento_medico.IdEstablecimiento)
	if error_create_paciente != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_paciente.Error(),
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
