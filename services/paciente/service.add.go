package paciente

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	paciente_repository "github.com/UPC-Works/api-aliviate/repositories/paciente"
)

func Add(c echo.Context) error {

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
	if len(input_paciente.Nombre) < 1 || len(input_paciente.Nombre) > 50 || len(input_paciente.Apellido) < 1 || len(input_paciente.Apellido) > 50 || input_paciente.Genero == 0 || input_paciente.DocumentoIdentidad == 0 || input_paciente.GrupoSanguineo == "" || input_paciente.RhSanguineo == "" || input_paciente.Telefono == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Nombre no puede exceder la longitud de 50, Apellido no puede exceder la longitud de 50, debe indicar el Genero y el Documento de Identidad",
			},
			Data: ""})
	}

	//Storage the Paciente
	var id_paciente string
	if input_paciente.Id == "" {
		id_paciente = uuid.New().String()
	} else {
		id_paciente = input_paciente.Id
	}
	new_paciente := models.NewPaciente(id_paciente, input_paciente.Nombre, input_paciente.Apellido, input_paciente.FechaNacimiento, input_paciente.Genero, input_paciente.DocumentoIdentidad, input_paciente.GrupoSanguineo, input_paciente.RhSanguineo, input_paciente.Telefono)
	error_create_paciente := paciente_repository.Pg_Create(new_paciente)
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
