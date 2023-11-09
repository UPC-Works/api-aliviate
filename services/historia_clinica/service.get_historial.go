package historia_clinica

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	modificaciones_historia_repository "github.com/UPC-Works/api-aliviate/repositories/modificaciones_historia"
)

func GetAllHistorial(c echo.Context) error {

	//Get the filters from the client
	idHistoriaClinica := c.Request().URL.Query().Get("idHistoriaClinica")

	//Get the all historial cambios
	list_historial, error_find_historial := modificaciones_historia_repository.Pg_FindMultiple(idHistoriaClinica)
	if error_find_historial != nil {
		return c.JSON(500, &helpers.ResponseListHistorialCambios{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_historial.Error(),
			},
			Data: []models.ModificacionesHistorias{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListHistorialCambios{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_historial})
}
