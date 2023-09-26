package consulta

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	consulta_repository "github.com/UPC-Works/api-aliviate/repositories/consulta"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	idHistoriaClinica := c.Request().URL.Query().Get("idHistoriaClinica")

	documentoIdentidad_string := c.Request().URL.Query().Get("documentoIdentidad")
	limit_string := c.Request().URL.Query().Get("limit")
	offset_string := c.Request().URL.Query().Get("offset")
	documentoIdentidad, _ := strconv.Atoi(documentoIdentidad_string)
	limit, _ := strconv.Atoi(limit_string)
	offset, _ := strconv.Atoi(offset_string)

	//Get the all consultas
	list_consultas, error_find_consultas := consulta_repository.Pg_FindMultiple(idHistoriaClinica, documentoIdentidad, limit, offset)
	if error_find_consultas != nil {
		return c.JSON(500, &helpers.ResponseListConsulta{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_consultas.Error(),
			},
			Data: []models.Consulta{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListConsulta{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_consultas})
}
