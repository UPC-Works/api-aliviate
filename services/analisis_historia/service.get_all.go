package analisis_historia

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	analisis_historia_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_historia"
)

func GetAll(c echo.Context) error {

	idAnalisisCodigo_string := c.Request().URL.Query().Get("idAnalisisCodigo")
	idAnalisisCodigo, _ := strconv.Atoi(idAnalisisCodigo_string)
	idHistoria := c.Request().URL.Query().Get("idHistoria")

	//Get all analisis_laboratorio
	valores_analisis_historia, _ := analisis_historia_repository.Pg_FindMultiple(idHistoria, idAnalisisCodigo)

	//OK
	return c.JSON(200, &helpers.ResponseListAnalisisHistoria{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: valores_analisis_historia})
}
