package establecimiento

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	establecimiento_repository "github.com/UPC-Works/api-aliviate/repositories/establecimiento"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_establecimiento *models.Establecimiento

	//Bind the model
	err := c.Bind(&input_establecimiento)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if len(input_establecimiento.Nombre) < 1 || len(input_establecimiento.Nombre) > 100 {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Nombre no puede exceder la longitud de 100",
			},
			Data: ""})
	}

	//Storage the Establecimiento
	new_establecimiento := models.NewEstablecimiento(uuid.New().String(), input_establecimiento.IdDistrito, input_establecimiento.Nombre, input_establecimiento.Dirección)
	error_create_establecimiento := establecimiento_repository.Pg_Create(new_establecimiento)
	if error_create_establecimiento != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_establecimiento.Error(),
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
