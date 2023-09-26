package analisis_laboratorio

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	analisis_laboratorio_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio"
)

func Update(c echo.Context) error {

	//Inicilization
	var input_analisis_laboratorio *models.AnalisisLaboratorio

	//Bind the model
	err := c.Bind(&input_analisis_laboratorio)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_analisis_laboratorio.Id == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Debe enviar el Id",
			},
			Data: ""})
	}

	//Find in the Storage
	analisis_laboratorio_found, error_find := analisis_laboratorio_repository.Pg_FindOne(input_analisis_laboratorio.IdHistoriaClinica)
	if error_find != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_find.Error(),
			},
			Data: ""})
	}
	if analisis_laboratorio_found.Id == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "No se ha encontrado el analisis de laboratorio",
			},
			Data: ""})
	}

	//Storage the AnalisisLaboratorio
	update_analisis_laboratorio := models.UpdateAnalisisLaboratorio(input_analisis_laboratorio.Id, input_analisis_laboratorio.IdHistoriaClinica, input_analisis_laboratorio.Colesterol, input_analisis_laboratorio.Trigliceridos, input_analisis_laboratorio.ColesterolHdl, input_analisis_laboratorio.ColesterolLdl, input_analisis_laboratorio.ColesterolVldl, input_analisis_laboratorio.Riesgo1, input_analisis_laboratorio.Riesgo2, input_analisis_laboratorio.Glucosa, input_analisis_laboratorio.Hematrocito, input_analisis_laboratorio.Hemoglobina)
	error_create_analisis_laboratorio := analisis_laboratorio_repository.Pg_Create(update_analisis_laboratorio)
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
