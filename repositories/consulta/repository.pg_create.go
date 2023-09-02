package consulta

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_consulta *models.Consulta) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Consulta (
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
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26)`
	_, err_query := db.Exec(ctx, query,
		input_consulta.Id,
		input_consulta.IdHistoriaClinica,
		input_consulta.FechaRegistro,
		input_consulta.DescripcionEnfermedadPaciente,
		input_consulta.TiempoEnfermedad,
		input_consulta.Apetito,
		input_consulta.Sed,
		input_consulta.Suenio,
		input_consulta.EstadoAnimo,
		input_consulta.OtroDetalle,
		input_consulta.Orina,
		input_consulta.Deposiciones,
		input_consulta.Temperatura,
		input_consulta.PA,
		input_consulta.FC,
		input_consulta.FR,
		input_consulta.Peso,
		input_consulta.Talla,
		input_consulta.IMC,
		input_consulta.Diagnostico,
		input_consulta.Tratamiento,
		input_consulta.DiagnosticoIA,
		input_consulta.TratamientoIA,
		input_consulta.ExamenesAuxiliares,
		input_consulta.ProximaCita,
		input_consulta.AtentidoPor,
		input_consulta.Observaciones,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
