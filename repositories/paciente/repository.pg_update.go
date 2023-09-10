package paciente

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_UpdateOne(input_paciente *models.Paciente) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `UPDATE Paciente SET
	nombre=$1            ,
	apellido=$2           ,
	fecha_nacimiento=$3   ,
	genero=$4             ,
	documento_identidad=$5,
	grupo_sanguineo=$6,
	rh_sanguineo=$7,
	telefono=$8 WHERE id=$9`
	_, err_query := db.Exec(ctx, query, input_paciente.Nombre, input_paciente.Apellido, input_paciente.FechaNacimiento, input_paciente.Genero, input_paciente.DocumentoIdentidad, input_paciente.GrupoSanguineo, input_paciente.RhSanguineo, input_paciente.Telefono, input_paciente.Id)

	if err_query != nil {
		return err_query
	}

	return nil
}
