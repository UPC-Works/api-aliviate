package prediccion

import (
	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	"github.com/labstack/echo/v4"
)

func Predecir(c echo.Context) error {

	predicciones := []models.PrediccionEnfermedad{
		{
			Enfermedad:   "Resfriado",
			Probabilidad: 0.75,
		},
		{
			Enfermedad:   "Gripe",
			Probabilidad: 0.95,
		},
		{
			Enfermedad:   "Dolor de garganta",
			Probabilidad: 0.60,
		},
	}

	//OK
	return c.JSON(200, &helpers.ResponsePrediccion{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: predicciones})
}
