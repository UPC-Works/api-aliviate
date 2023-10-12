package analisis_laboratorio

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	analisis_laboratorio_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio"
)

func GetOne(c echo.Context) error {

	//Get the id analisis laboratorio
	id_historia_clinica := c.Param("id_historia_clinica")

	//Get one analisis_laboratorio
	analisis_laboratorio, _ := analisis_laboratorio_repository.Pg_FindOne(id_historia_clinica)

	//OK
	return c.JSON(200, &helpers.ResponseAnalisisLaboratorio{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: analisis_laboratorio})
}
