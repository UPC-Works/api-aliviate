package paciente

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_paciente *models.Paciente) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Paciente (
		id                 ,
		nombre             ,
		apellido           ,
		fecha_nacimiento   ,
		genero             ,
		documento_identidad,
		fecha_registro     ,
		grupo_sanguineo,
		rh_sanguineo,
		telefono
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`
	_, err_query := db.Exec(ctx, query,
		input_paciente.Id,
		input_paciente.Nombre,
		input_paciente.Apellido,
		input_paciente.FechaNacimiento,
		input_paciente.Genero,
		input_paciente.DocumentoIdentidad,
		input_paciente.FechaRegistro,
		input_paciente.GrupoSanguineo,
		input_paciente.RhSanguineo,
		input_paciente.Telefono,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
