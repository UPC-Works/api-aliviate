package establecimiento_medico

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
)

func Pg_Create(input_idmedico string, input_idestablecimiento string) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO EstablecimientoMedico (
		id_establecimiento                 ,
		id_medico
	) VALUES ($1,$2)`
	_, err_query := db.Exec(ctx, query,
		input_idestablecimiento,
		input_idmedico,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
