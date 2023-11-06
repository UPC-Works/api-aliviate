package analisis_historia

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	analisis_historia_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_historia"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_analisis_historia []*models.AnalisisHistoria

	//Bind the model
	err := c.Bind(&input_analisis_historia)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Check the structure or the type of the value",
			},
			Data: ""})
	}

	//Storage the Analisis de Historia
	error_create_analisis_historia := analisis_historia_repository.Pg_Create(input_analisis_historia)
	if error_create_analisis_historia != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_analisis_historia.Error(),
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
