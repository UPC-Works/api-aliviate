package analisis_laboratorio

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	analisis_laboratorio_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_analisis_laboratorio *models.AnalisisLaboratorio

	//Bind the model
	err := c.Bind(&input_analisis_laboratorio)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Check the structure or the type of the value",
			},
			Data: ""})
	}

	//Storage the Analisis de laboratorio
	new_analisis_laboratorio := models.NewAnalisisLaboratorio(uuid.New().String(), input_analisis_laboratorio.IdHistoriaClinica, input_analisis_laboratorio.Colesterol, input_analisis_laboratorio.Trigliceridos, input_analisis_laboratorio.ColesterolHdl, input_analisis_laboratorio.ColesterolLdl, input_analisis_laboratorio.ColesterolVldl, input_analisis_laboratorio.Riesgo1, input_analisis_laboratorio.Riesgo2, input_analisis_laboratorio.Glucosa, input_analisis_laboratorio.Hematrocito, input_analisis_laboratorio.Hemoglobina)
	error_create_analisis_laboratorio := analisis_laboratorio_repository.Pg_Create(new_analisis_laboratorio)
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
