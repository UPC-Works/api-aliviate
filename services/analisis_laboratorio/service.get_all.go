package analisis_laboratorio

import (
	"strconv"

	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	analisis_laboratorio_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio"
)

func GetAll(c echo.Context) error {

	//Get the filters from the client
	idHistoriaClinica := c.Request().URL.Query().Get("idHistoriaClinica")
	limit_string := c.Request().URL.Query().Get("limit")
	offset_string := c.Request().URL.Query().Get("offset")
	limit, _ := strconv.Atoi(limit_string)
	offset, _ := strconv.Atoi(offset_string)

	//Get all analisis_laboratorio
	analisis_laboratorio, _ := analisis_laboratorio_repository.Pg_FindMultiple(idHistoriaClinica, limit, offset)

	//OK
	return c.JSON(200, &helpers.ResponseListAnalisisLaboratorio{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: analisis_laboratorio})
}
