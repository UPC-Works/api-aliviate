package establecimiento

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	consulta_repository "github.com/UPC-Works/api-aliviate/repositories/consulta"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_consulta *models.Consulta

	//Bind the model
	err := c.Bind(&input_consulta)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_consulta.IdHistoriaClinica == "" || input_consulta.IdMedico == "" {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "IdHistoriaClinica y IdMedico deben ser enviados",
			},
			Data: ""})
	}

	//Storage the Consulta
	new_consulta := models.NewConsulta(uuid.New().String(), input_consulta.IdHistoriaClinica, input_consulta.IdMedico, input_consulta.DescripcionEnfermedadPaciente, input_consulta.TiempoEnfermedad, input_consulta.Apetito, input_consulta.Sed, input_consulta.Suenio, input_consulta.EstadoAnimo, input_consulta.OtroDetalle, input_consulta.Orina, input_consulta.Deposiciones, input_consulta.Temperatura, input_consulta.PA, input_consulta.FC, input_consulta.FR, input_consulta.Peso, input_consulta.Talla, input_consulta.IMC, input_consulta.Diagnostico, input_consulta.Tratamiento, input_consulta.DiagnosticoIA, input_consulta.TratamientoIA, input_consulta.ExamenesAuxiliares, input_consulta.ProximaCita, input_consulta.Observaciones)
	error_create_consulta := consulta_repository.Pg_Create(new_consulta)
	if error_create_consulta != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_consulta.Error(),
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
