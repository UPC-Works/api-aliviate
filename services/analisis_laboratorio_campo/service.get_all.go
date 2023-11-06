package analisis_laboratorio

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	analisis_laboratorio_campo_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio_campo"
)

func GetAll(c echo.Context) error {

	idAnalisisCodigo_string := c.Request().URL.Query().Get("idAnalisisCodigo")
	idAnalisisCodigo, _ := strconv.Atoi(idAnalisisCodigo_string)

	//Get all analisis_laboratorio
	campos_analisis_laboratorio, _ := analisis_laboratorio_campo_repository.Pg_FindMultiple(idAnalisisCodigo)

	//OK
	return c.JSON(200, &helpers.ResponseListAnalisisLaboratorioCampo{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: campos_analisis_laboratorio})
}
