package estadistica

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	diagnostico_ia_repository "github.com/UPC-Works/api-aliviate/repositories/diagnostico_ia"
)

func GetEnfermedadProbabilidad(c echo.Context) error {

	//Get the id form the Middleware
	idMedico := c.Get("id").(string)

	//Get the all probabilidad_enfermedades
	list_probabilidad_enfermedades, error_find_probabilidad_enfermedades := diagnostico_ia_repository.Pg_Find_EnfermedadesPredicciones(idMedico)
	if error_find_probabilidad_enfermedades != nil {
		return c.JSON(500, &helpers.ResponseListEnfermedadPrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find_probabilidad_enfermedades.Error(),
			},
			Data: []models.EnfermedadPrediccion{}})
	}

	//OK
	return c.JSON(200, &helpers.ResponseListEnfermedadPrediccion{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: list_probabilidad_enfermedades})
}
