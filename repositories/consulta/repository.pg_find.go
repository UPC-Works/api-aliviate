package consulta

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindMultiple(input_idhistoriaclinica string, input_limit int, input_offset int) ([]models.Consulta, error) {

	//Initialization
	var oListConsulta []models.Consulta

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idhistoriaclinica != "" {
		filters["id_historia_clinica"] = input_idhistoriaclinica
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := configs.Conn_Pg_DB()

	//Define the query
	q := `SELECT 
	id                             ,
	id_historia_clinica            ,
	fecha_registro                 ,
	descripcion_enfermedad_paciente,
	tiempo_enfermedad              ,
	apetito                        ,
	sed                            ,
	suenio                         ,
	estado_anio                    ,
	otro_detalle                   ,
	orina                          ,
	deposiciones                   ,
	temperatura                    ,
	p_a                            ,
	f_c                            ,
	f_r                            ,
	peso                           ,
	talla                          ,
	i_m_c                          ,
	diagnostico                    ,
	tratamiento                    ,
	diagnostico_ia                 ,
	tratamiento_ia                 ,
	tratamiento                    ,
	examenes_auxiliares            ,
	proxima_cita                   ,
	atendido_por                   ,
	observaciones                  
FROM Consulta `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" ORDER BY fecha_registro DESC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListConsulta, error_find
	}

	//Scan the row
	for rows.Next() {
		var oConsulta models.Consulta
		rows.Scan(
			&oConsulta.Id,
			&oConsulta.IdHistoriaClinica,
			&oConsulta.FechaRegistro,
			&oConsulta.DescripcionEnfermedadPaciente,
			&oConsulta.TiempoEnfermedad,
			&oConsulta.Apetito,
			&oConsulta.Sed,
			&oConsulta.Suenio,
			&oConsulta.EstadoAnimo,
			&oConsulta.OtroDetalle,
			&oConsulta.Orina,
			&oConsulta.Deposiciones,
			&oConsulta.Temperatura,
			&oConsulta.PA,
			&oConsulta.FC,
			&oConsulta.FR,
			&oConsulta.Peso,
			&oConsulta.Talla,
			&oConsulta.IMC,
			&oConsulta.Diagnostico,
			&oConsulta.Tratamiento,
			&oConsulta.DiagnosticoIA,
			&oConsulta.TratamientoIA,
			&oConsulta.ExamenesAuxiliares,
			&oConsulta.ProximaCita,
			&oConsulta.AtentidoPor,
			&oConsulta.Observaciones)
		oListConsulta = append(oListConsulta, oConsulta)
	}

	if error_find != nil {
		return oListConsulta, error_find
	}

	//Return the list of consultas
	return oListConsulta, nil
}
