package paciente

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	paciente_repository "github.com/UPC-Works/api-aliviate/repositories/paciente"
)

func Update(c echo.Context) error {

	//Inicilization
	var input_paciente *models.Paciente

	//Bind the model
	err := c.Bind(&input_paciente)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_paciente.Id == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Debe enviar el Id",
			},
			Data: ""})
	}

	//Find in the Storage
	paciente_found, error_find := paciente_repository.Pg_FindOne(input_paciente.DocumentoIdentidad)
	if error_find != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find.Error(),
			},
			Data: ""})
	}
	if paciente_found.Id == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "No se ha encontrado el paciente",
			},
			Data: ""})
	}

	//Storage the Paciente
	update_paciente := models.UpdatePaciente(input_paciente.Id, input_paciente.Nombre, input_paciente.Apellido, input_paciente.FechaNacimiento, input_paciente.Genero, input_paciente.DocumentoIdentidad)
	error_create_paciente := paciente_repository.Pg_Create(update_paciente)
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
