package establecimiento

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_establecimiento *models.Establecimiento) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Establecimiento (
		id                 ,
		nombre             
	) VALUES ($1,$2)`
	_, err_query := db.Exec(ctx, query,
		input_establecimiento.Id,
		input_establecimiento.Nombre,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
