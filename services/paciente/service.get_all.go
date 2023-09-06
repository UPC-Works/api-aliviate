package paciente

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	paciente_repository "github.com/UPC-Works/api-aliviate/repositories/paciente"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	idPaciente := c.Request().URL.Query().Get("idPaciente")
	dni_string := c.Request().URL.Query().Get("dni")
	limit_string := c.Request().URL.Query().Get("limit")
	offset_string := c.Request().URL.Query().Get("offset")
	dni, _ := strconv.Atoi(dni_string)
	limit, _ := strconv.Atoi(limit_string)
	offset, _ := strconv.Atoi(offset_string)

	//Get the all pacientes
	list_pacientes, error_find_pacientes := paciente_repository.Pg_FindMultiple(idPaciente, dni, limit, offset)
	if error_find_pacientes != nil {
		return c.JSON(500, &helpers.ResponseListPaciente{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_pacientes.Error(),
			},
			Data: []models.Paciente{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListPaciente{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_pacientes})
}
