package medico

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	medico_repository "github.com/UPC-Works/api-aliviate/repositories/medico"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	name := c.Request().URL.Query().Get("name")

	//Get the all medicos
	list_medicos, error_find_medicos := medico_repository.Pg_FindMultiple(name)
	if error_find_medicos != nil {
		return c.JSON(500, &helpers.ResponseListMedico{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_medicos.Error(),
			},
			Data: []models.Medico{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListMedico{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_medicos})
}
