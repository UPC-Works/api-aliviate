package establecimiento

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	establecimiento_repository "github.com/UPC-Works/api-aliviate/repositories/establecimiento"
)

func GetAll(c echo.Context) error {

	//Get the all establecimientos
	list_establecimientos, error_find_establecimientos := establecimiento_repository.Pg_FindMultiple()
	if error_find_establecimientos != nil {
		return c.JSON(500, &helpers.ResponseListEstablecimiento{
			Error: helpers.ErrorStructure{
				Code:   9457,
				Detail: error_find_establecimientos.Error(),
			},
			Data: []models.Establecimiento{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListEstablecimiento{
		Error: helpers.ErrorStructure{
			Code:   0,
			Detail: "",
		},
		Data: list_establecimientos})
}
