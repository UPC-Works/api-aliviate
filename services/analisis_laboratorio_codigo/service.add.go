package analisis_laboratorio

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	analisis_laboratorio_codigo_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio_codigo"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_analisis_laboratorio_codigo *models.AnalisisLaboratorioCodigo

	//Bind the model
	err := c.Bind(&input_analisis_laboratorio_codigo)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Check the structure or the type of the value",
			},
			Data: ""})
	}

	//Storage the Analisis de laboratorio codigo
	new_analisis_laboratorio := models.NewAnalisisLaboratorio(input_analisis_laboratorio_codigo.Id, input_analisis_laboratorio_codigo.Nombre)
	error_create_analisis_laboratorio := analisis_laboratorio_codigo_repository.Pg_Create(new_analisis_laboratorio)
	if error_create_analisis_laboratorio != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_analisis_laboratorio.Error(),
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
