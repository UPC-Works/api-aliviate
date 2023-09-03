package admin

import (
	"github.com/labstack/echo/v4"
	"github.com/lithammer/shortuuid/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	admin_repository "github.com/UPC-Works/api-aliviate/repositories/admin"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func SignUp(c echo.Context) error {

	//Inicilization
	var input_admin *models.Admin

	//Bind the model
	err := c.Bind(&input_admin)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9451,
				Detail: "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if len(input_admin.NombreCompleto) < 1 || len(input_admin.NombreCompleto) > 50 || len(input_admin.Correo) < 1 || len(input_admin.Correo) > 50 || len(input_admin.Contrasenia) < 8 || len(input_admin.Contrasenia) > 12 {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9452,
				Detail: "El nombre no debe exceder la longitud de 20 caracteres, Correo no puede exceder la longitud de 50 y la Contraseña debe estar entre 8-12 caracteres",
			},
			Data: ""})
	}

	//Validation if the email already exist
	admin_found, _ := admin_repository.Pg_FindOne(input_admin.Correo)
	if admin_found.Id != "" {
		return c.JSON(403, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9456,
				Detail: "Este correo ya esta registrado",
			},
			Data: ""})
	}

	//Encrypt password
	encrypted_pass, _ := private_services.EncryptPassword(admin_found.Contrasenia)

	//Storage the New Admin
	new_admin := models.NewAdmin(shortuuid.New(), admin_found.NombreCompleto, admin_found.Correo, encrypted_pass)
	error_create_admin := admin_repository.Pg_Create(new_admin)
	if error_create_admin != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9457,
				Detail: error_create_admin.Error(),
			},
			Data: ""})
	}

	//Storage in cache
	err_set := admin_repository.Re_SetId(new_admin.Id)
	if err_set != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9453,
				Detail: err_set.Error(),
			},
			Data: ""})
	}

	//OK
	return c.JSON(200, &helpers.ResponseString{
		Error: helpers.ErrorStructure{
			Code:   0,
			Detail: "",
		},
		Data: "OK"})
}
