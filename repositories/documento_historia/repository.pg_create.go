package documento_historia

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_documento_historia *models.DocumentosHistoria) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO DocumentoHistoria (
		id,
		id_historia_clinica,
		url
	) VALUES ($1,$2,$3)`
	_, err_query := db.Exec(ctx, query,
		input_documento_historia.Id,
		input_documento_historia.IdHistoriaClinica,
		input_documento_historia.Url,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
