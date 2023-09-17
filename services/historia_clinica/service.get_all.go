package historia_clinica

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	historia_clinica_repository "github.com/UPC-Works/api-aliviate/repositories/historia_clinica"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	documentoIdentidadPaciente_string := c.Request().URL.Query().Get("documentoIdentidadPaciente")
	limit_string := c.Request().URL.Query().Get("limit")
	offset_string := c.Request().URL.Query().Get("offset")
	documentoIdentidadPaciente, _ := strconv.Atoi(documentoIdentidadPaciente_string)
	limit, _ := strconv.Atoi(limit_string)
	offset, _ := strconv.Atoi(offset_string)

	//Get the all historia_clinicas
	list_historia_clinicas, error_find_historia_clinicas := historia_clinica_repository.Pg_FindMultiple(documentoIdentidadPaciente, limit, offset)
	if error_find_historia_clinicas != nil {
		return c.JSON(500, &helpers.ResponseListMedico{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_historia_clinicas.Error(),
			},
			Data: []models.Medico{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListHistoriaClinica{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_historia_clinicas})
}
