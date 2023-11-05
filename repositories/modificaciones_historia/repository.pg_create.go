package modificaciones_historia

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_modificaciones_historia *models.ModificacionesHistorias) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO ModificacionesHistoria (
		id                 ,
		id_historia_clinica             ,
		nombre_medico           ,
		actualizado_el       
	) VALUES ($1,$2,$3,$4)`
	_, err_query := db.Exec(ctx, query,
		input_modificaciones_historia.Id,
		input_modificaciones_historia.IdHistoriaClinica,
		input_modificaciones_historia.NombreMedico,
		input_modificaciones_historia.ActualizadoEl,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
