package historia_clinica

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	historia_clinica_repository "github.com/UPC-Works/api-aliviate/repositories/historia_clinica"
)

func GetOne(c echo.Context) error {

	//Get the filters from the client
	idHistoriaClinica := c.Request().URL.Query().Get("idHistoriaClinica")

	//Get one historia_clinicas
	historia_clinica, error_find_historia_clinica := historia_clinica_repository.Pg_FindOne(idHistoriaClinica)
	if error_find_historia_clinica != nil {
		return c.JSON(500, &helpers.ResponseHistoriaClinica{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_historia_clinica.Error(),
			},
			Data: historia_clinica})
	}

	//OK
	return c.JSON(200, &helpers.ResponseHistoriaClinica{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: historia_clinica})
}
