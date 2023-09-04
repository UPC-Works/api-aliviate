package medico

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_medico *models.Medico) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Medico (
		id                 ,
		nombre             ,
		apellido           ,
		colegiatura        ,
		correo             ,
		contrasenia        ,
		direccion          ,
		fecha_registro     ,
		especialidad
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`
	_, err_query := db.Exec(ctx, query,
		input_medico.Id,
		input_medico.Nombre,
		input_medico.Apellido,
		input_medico.Colegiatura,
		input_medico.Correo,
		input_medico.Contrasenia,
		input_medico.Direccion,
		input_medico.FechaRegistro,
		input_medico.Especialidad,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
