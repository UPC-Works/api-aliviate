package auth

import (
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	admin_repository "github.com/UPC-Works/api-aliviate/repositories/admin"
	medico_repository "github.com/UPC-Works/api-aliviate/repositories/medico"
	private_services "github.com/UPC-Works/api-aliviate/services/private"
)

func Authentication(c echo.Context) error {

	//Get the auth-token
	input_jwt := c.Request().Header.Get("auth-token")

	//Initialize claims
	claims := &models.Claim{}

	//Bind the model
	err := c.Bind(&input_jwt)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				Code:   9451,
				Detail: "Revisa la estructura o el tipo del valor, detalles: " + err.Error(),
			},
			Data: ""})
	}

	//Descrypt the JWT
	token, error_decrypt := private_services.DecryptJWT(input_jwt, claims)
	if error_decrypt != nil {
		return c.JSON(500, &helpers.ResponseAuth{
			Error: helpers.ErrorStructure{
				Code:   94510,
				Detail: "Error al desencriptar token,detalle:" + error_decrypt.Error(),
			},
			Data: helpers.AuthStructure{
				Id:             claims.Id,
				NombreCompleto: claims.NombreCompleto,
				Rol:            claims.Rol,
				Correo:         claims.Correo,
			}})
	}
	if !token.Valid {
		return c.JSON(400, &helpers.ResponseAuth{
			Error: helpers.ErrorStructure{
				Code:   9451,
				Detail: "Token invalido",
			},
			Data: helpers.AuthStructure{
				Id:             claims.Id,
				NombreCompleto: claims.NombreCompleto,
				Rol:            claims.Rol,
				Correo:         claims.Correo,
			}})
	}

	//If the session is from the Admin
	if claims.Rol == 1 {

		id_admin, error_get_re := admin_repository.Re_GetId(claims.Id)

		if id_admin == "" {
			return c.JSON(403, &helpers.ResponseAuth{
				Error: helpers.ErrorStructure{
					Code:   9459,
					Detail: "Restart your session",
				},
				Data: helpers.AuthStructure{
					Id:             claims.Id,
					NombreCompleto: claims.NombreCompleto,
					Rol:            claims.Rol,
					Correo:         claims.Correo,
				}})
		}

		if error_get_re != nil {
			//Get the last admin data updated
			admin_updated, error_find_admin := admin_repository.Pg_FindOne(claims.Correo)
			if error_find_admin != nil {
				return c.JSON(500, &helpers.ResponseAuth{
					Error: helpers.ErrorStructure{
						Code:   9457,
						Detail: error_find_admin.Error(),
					},
					Data: helpers.AuthStructure{
						Id:             claims.Id,
						NombreCompleto: claims.NombreCompleto,
						Rol:            claims.Rol,
						Correo:         claims.Correo,
					}})
			}
			//Storage the session_code in the cache
			err_add_re := admin_repository.Re_SetId(admin_updated.Id)
			if err_add_re != nil {
				return c.JSON(500, &helpers.ResponseAuth{
					Error: helpers.ErrorStructure{
						Code:   9453,
						Detail: err_add_re.Error(),
					},
					Data: helpers.AuthStructure{
						Id:             claims.Id,
						NombreCompleto: claims.NombreCompleto,
						Rol:            claims.Rol,
						Correo:         claims.Correo,
					}})
			}
			//Return the error
			return c.JSON(500, &helpers.ResponseAuth{
				Error: helpers.ErrorStructure{
					Code:   9453,
					Detail: error_get_re.Error(),
				},
				Data: helpers.AuthStructure{
					Id:             claims.Id,
					NombreCompleto: claims.NombreCompleto,
					Rol:            claims.Rol,
					Correo:         claims.Correo,
				}})
		}
	} else {

		id_medico, error_get_re := admin_repository.Re_GetId(claims.Id)

		if id_medico == "" {
			return c.JSON(403, &helpers.ResponseAuth{
				Error: helpers.ErrorStructure{
					Code:   9459,
					Detail: "Restablece tu sesion",
				},
				Data: helpers.AuthStructure{
					Id:             claims.Id,
					NombreCompleto: claims.NombreCompleto,
					Rol:            claims.Rol,
					Correo:         claims.Correo,
				}})
		}

		if error_get_re != nil {
			//Get the last medico data updated
			medico_updated, error_find_medico := medico_repository.Pg_FindOne(claims.Correo)
			if error_find_medico != nil {
				return c.JSON(500, &helpers.ResponseAuth{
					Error: helpers.ErrorStructure{
						Code:   9457,
						Detail: error_find_medico.Error(),
					},
					Data: helpers.AuthStructure{
						Id:             claims.Id,
						NombreCompleto: claims.NombreCompleto,
						Rol:            claims.Rol,
						Correo:         claims.Correo,
					}})
			}
			//Storage in the cache
			err_add_re := medico_repository.Re_SetId(medico_updated.Id)
			if err_add_re != nil {
				return c.JSON(500, &helpers.ResponseAuth{
					Error: helpers.ErrorStructure{
						Code:   9453,
						Detail: err_add_re.Error(),
					},
					Data: helpers.AuthStructure{
						Id:             claims.Id,
						NombreCompleto: claims.NombreCompleto,
						Rol:            claims.Rol,
						Correo:         claims.Correo,
					}})
			}
			//Return the error
			return c.JSON(500, &helpers.ResponseAuth{
				Error: helpers.ErrorStructure{
					Code:   9453,
					Detail: error_get_re.Error(),
				},
				Data: helpers.AuthStructure{
					Id:             claims.Id,
					NombreCompleto: claims.NombreCompleto,
					Rol:            claims.Rol,
					Correo:         claims.Correo,
				}})
		}
	}

	//OK
	return c.JSON(200, &helpers.ResponseAuth{
		Error: helpers.ErrorStructure{
			Code:   0,
			Detail: "",
		},
		Data: helpers.AuthStructure{
			Id:             claims.Id,
			NombreCompleto: claims.NombreCompleto,
			Rol:            claims.Rol,
			Correo:         claims.Correo,
		}})
}
