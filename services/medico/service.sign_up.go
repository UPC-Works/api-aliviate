package medico

import (
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	medico_repository "github.com/UPC-Works/api-aliviate/repositories/medico"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func SignUp(c echo.Context) error {

	//Inicilization
	var input_medico *models.Medico

	//Bind the model
	err := c.Bind(&input_medico)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if len(input_medico.Nombre) < 1 || len(input_medico.Nombre) > 20 || len(input_medico.Correo) < 1 || len(input_medico.Correo) > 50 || len(input_medico.Contrasenia) < 8 || len(input_medico.Contrasenia) > 12 {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "El nombre no debe exceder la longitud de 20 caracteres, Correo no puede exceder la longitud de 50 y la Contraseña debe estar entre 8-12 caracteres",
			},
			Data: ""})
	}

	//Validation if the email already exist
	medico_found, _ := medico_repository.Pg_FindOne(input_medico.Correo)
	if medico_found.Id != "" {
		return c.JSON(403, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Este correo ya esta registrado",
			},
			Data: ""})
	}

	//Encrypt password
	encrypted_pass, _ := private_services.EncryptPassword(input_medico.Contrasenia)

	//Storage the New Medico
	new_medico := models.NewMedico(shortuuid.New(), input_medico.Nombre, input_medico.Apellido, input_medico.Colegiatura, input_medico.DocumentoIdentidad, input_medico.Correo, encrypted_pass, input_medico.Direccion, input_medico.Especialidad)
	error_create_medico := medico_repository.Pg_Create(new_medico)
	if error_create_medico != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_medico.Error(),
			},
			Data: ""})
	}

	//Storage in the cache
	err_set := medico_repository.Re_SetId(new_medico.Id)
	if err_set != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   err_set.Error(),
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
