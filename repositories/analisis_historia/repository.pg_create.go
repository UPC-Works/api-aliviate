package analisis_historia

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_analisis_historias []*models.AnalisisHistoria) error {

	//Variables for bulk insert
	id_historia_clinica_pg, id_analisis_campo_pg, valor_pg := []string{}, []int{}, []float32{}

	for _, input := range input_analisis_historias {
		id_historia_clinica_pg = append(id_historia_clinica_pg, input.IdHistoriaClinica)
		id_analisis_campo_pg = append(id_analisis_campo_pg, input.IdAnalisisCampo)
		valor_pg = append(valor_pg, input.Valor)
	}

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO AnalisisHistoria (id_historia_clinica,id_analisis_campo,valor) VALUES (select * from 
		unnest($1::VARCHAR(50)[],$2::serial[], $3::valor[]))`
	_, err_query := db.Exec(ctx, query, id_historia_clinica_pg, id_analisis_campo_pg, valor_pg)

	if err_query != nil {
		return err_query
	}

	return nil
}
