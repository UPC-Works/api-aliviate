package establecimiento

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
	"github.com/google/uuid"
)

func Pg_Create(input_consulta *models.Consulta) error {

	//Variables for bulk insert
	id_pg, id_consulta_pg, enfermedad_pg, probabilidad_pg, es_aceptado_pg := []string{}, []string{}, []string{}, []float32{}, []bool{}

	for _, input_prediccion := range input_consulta.DiagnosticoIA {
		id_pg = append(id_pg, uuid.New().String())
		id_consulta_pg = append(id_consulta_pg, input_consulta.Id)
		enfermedad_pg = append(enfermedad_pg, input_prediccion.Enfermedad)
		probabilidad_pg = append(probabilidad_pg, input_prediccion.Probabilidad)
		es_aceptado_pg = append(es_aceptado_pg, input_prediccion.EsAceptado)
	}

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO DiagnosticoIA (id,id_consulta,enfermedad,probabilidad,es_aceptado) VALUES (select * from 
		unnest($1::VARCHAR(50)[], $2::VARCHAR(50)[],$3::VARCHAR(50)[],$4::decimal(10,2)[],$5::bool[]))`
	_, err_query := db.Exec(ctx, query, id_pg, id_pg, id_consulta_pg, enfermedad_pg, probabilidad_pg, es_aceptado_pg)

	if err_query != nil {
		return err_query
	}

	return nil
}
