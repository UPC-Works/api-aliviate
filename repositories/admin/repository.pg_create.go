package admin

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_admin *models.Admin) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Admin (
		id                 ,
		nombre_completo             .
		correo,
		contrasenia
	) VALUES ($1,$2,$3,$4)`
	_, err_query := db.Exec(ctx, query,
		input_admin.Id,
		input_admin.NombreCompleto,
		input_admin.Correo,
		input_admin.Contrasenia,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
