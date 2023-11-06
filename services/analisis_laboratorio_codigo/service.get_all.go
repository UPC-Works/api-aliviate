package analisis_laboratorio

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	analisis_laboratorio_codigo_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio_codigo"
)

func GetAll(c echo.Context) error {

	//Get all analisis_laboratorio
	codigos_analisis_laboratorio, _ := analisis_laboratorio_codigo_repository.Pg_FindMultiple()

	//OK
	return c.JSON(200, &helpers.ResponseListAnalisisLaboratorioCodigo{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: codigos_analisis_laboratorio})
}
