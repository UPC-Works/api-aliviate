package analisis_laboratorio_campo

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_analisis_lab_campos []*models.AnalisisLaboratorioCampo) error {

	//Variables for bulk insert
	id_analisis_pg, campo_pg, campo_json_pg, tipo_pg := []int{}, []string{}, []string{}, []string{}

	for _, input := range input_analisis_lab_campos {
		id_analisis_pg = append(id_analisis_pg, input.IdAnalisis)
		campo_pg = append(campo_pg, input.Campo)
		campo_json_pg = append(campo_json_pg, input.CampoJson)
		tipo_pg = append(tipo_pg, input.Tipo)
	}

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO AnalisisLaboratorioCampo (id_analisis,campo,campos_json,tipo) VALUES (select * from 
		unnest($1::integer[],$2::VARCHAR(50)[], $3::VARCHAR(50)[], $4::VARCHAR(50)[]))`
	_, err_query := db.Exec(ctx, query, id_analisis_pg, campo_pg, campo_json_pg, tipo_pg)

	if err_query != nil {
		return err_query
	}

	return nil
}
