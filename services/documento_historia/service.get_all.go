package documento_historia

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	documento_historia_repository "github.com/UPC-Works/api-aliviate/repositories/documento_historia"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	idHistoriaClinica := c.Request().URL.Query().Get("idHistoriaClinica")

	//Get the all documentos
	list_documentos, error_find_documentos := documento_historia_repository.Pg_FindMultiple(idHistoriaClinica)
	if error_find_documentos != nil {
		return c.JSON(500, &helpers.ResponseListMedico{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_documentos.Error(),
			},
			Data: []models.Medico{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListDocumentos{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_documentos})
}
