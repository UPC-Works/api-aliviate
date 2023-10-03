package consulta

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_idhistoriaclinica string) (models.Consulta, error) {

	//Initialization
	var oConsulta models.Consulta

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idhistoriaclinica != "" {
		filters["co.id_historia_clinica"] = input_idhistoriaclinica
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
	co.id                             ,
	co.id_historia_clinica            ,
	co.fecha_registro                 ,
	co.descripcion_enfermedad_paciente,
	co.tiempo_enfermedad              ,
	co.apetito                        ,
	co.sed                            ,
	co.suenio                         ,
	co.estado_animo                    ,
	co.otro_detalle                   ,
	co.orina                          ,
	co.deposiciones                   ,
	co.temperatura                    ,
	co.p_a                            ,
	co.f_c                            ,
	co.f_r                            ,
	co.peso                           ,
	co.talla                          ,
	co.i_m_c                          ,
	co.diagnostico                    ,
	co.tratamiento                    ,
	co.diagnostico_ia                 ,
	co.tratamiento_ia                 ,
	co.tratamiento                    ,
	co.examenes_auxiliares            ,
	co.proxima_cita                   ,
	co.id_medico                   ,
	co.observaciones                  ,
	concat(pa.nombre,' ',pa.apellido),
	pa.id
	FROM Consulta AS co JOIN HistoriaClinica AS hc ON co.id_historia_clinica=hc.id JOIN Paciente AS pa ON hc.id_paciente=pa.id`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
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
		&oConsulta.IdMedico,
		&oConsulta.Observaciones,
		&oConsulta.NombreCompletoPaciente,
		&oConsulta.IdPaciente)

	if error_find != nil {
		return oConsulta, error_find
	}

	//Return one consulta
	return oConsulta, nil
}

func Pg_FindMultiple(input_idhistoriaclinica string, input_documento_identidad int, input_limit int, input_offset int) ([]models.Consulta, error) {

	//Initialization
	var oListConsulta []models.Consulta

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_idhistoriaclinica != "" {
		filters["co.id_historia_clinica"] = input_idhistoriaclinica
		counter_filters += 1
	}
	if input_documento_identidad != 0 {
		filters["pa.documento_identidad"] = input_documento_identidad
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
	co.id                             ,
	co.id_historia_clinica            ,
	co.fecha_registro                 ,
	co.descripcion_enfermedad_paciente,
	co.tiempo_enfermedad              ,
	co.apetito                        ,
	co.sed                            ,
	co.suenio                         ,
	co.estado_animo                    ,
	co.otro_detalle                   ,
	co.orina                          ,
	co.deposiciones                   ,
	co.temperatura                    ,
	co.p_a                            ,
	co.f_c                            ,
	co.f_r                            ,
	co.peso                           ,
	co.talla                          ,
	co.i_m_c                          ,
	co.diagnostico                    ,
	co.tratamiento                    ,
	co.diagnostico_ia                 ,
	co.tratamiento_ia                 ,
	co.examenes_auxiliares            ,
	co.proxima_cita                   ,
	co.id_medico                   ,
	co.observaciones                  ,
	concat(pa.nombre,' ',pa.apellido),
	pa.id
FROM Consulta AS co JOIN HistoriaClinica AS hc ON co.id_historia_clinica=hc.id JOIN Paciente AS pa ON hc.id_paciente=pa.id`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			if key == "co.documento_identidad" {
				clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
			} else {
				clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
			}
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
			&oConsulta.IdMedico,
			&oConsulta.Observaciones,
			&oConsulta.NombreCompletoPaciente,
			&oConsulta.IdPaciente)
		oListConsulta = append(oListConsulta, oConsulta)
	}

	if error_find != nil {
		return oListConsulta, error_find
	}

	//Return the list of consultas
	return oListConsulta, nil
}
