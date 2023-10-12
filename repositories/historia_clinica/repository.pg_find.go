package historia_clinica

import (
	"context"
	"fmt"
	"strings"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_FindOne(input_id string) (models.HistoriaClinica, error) {

	//Initialization
	var oHistoriaClinica models.HistoriaClinica

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_id != "" {
		filters["id"] = input_id
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
	id                                ,
		id_medico,
		id_paciente,
		estado_civil,
		grado_institucion,
		ocupacion,
		direccion,
		id_distrito                       ,
		fecha_registro                    ,
		tuvo_tuberculosis                 ,
		tiene_inf_trans_sex               ,
		tiene_diabetes                    ,
		tiene_hta                         ,
		tiene_sobrepeso                   ,
		tiene_infarto                     ,
		tiene_dislipenia                  ,
		tiene_inf_renal_glaucoma          ,
		tiene_depresion_esquizofrenia     ,
		antecedentes                      ,
		tiene_hospitaliacion_transfusiones,
		dispacidad                        ,
		tiene_consumo_tabaco              ,
		tiene_consumo_alcohol             ,
		tiene_consumo_drogas              ,
		tiene_inter_quirurjica            ,
		cancer                            ,
		tiene_riesgo                      ,
		tiene_violencia                   ,
		tiene_sida                        ,
		tiene_its                         ,
		tiene_hepatitis                   ,
		tiene_dbm                         ,
		tiene_cancer                      ,
		tiene_depresion                   ,
		tiene_prob_psiquiatricos          ,
		otros                             ,
		reaccion_medicamentos             ,
		medicamente_frecuente             ,
		edad_inicio_relacion_sexual       ,
		num_parejas                       ,
		hijos_vivos                       ,
		rs_mismo_sexo                     ,
		menarquia                         ,
		flujo_vag_patologico              ,
		dismenorrea                       ,
		tiene_embarazo                    ,
		tiene_parto                       ,
		tiene_prematuro                   ,
		tiene_aborto                      ,
		gestacion                         ,
		tiene_fiebre_15_dias              ,
		tiene_tos_15_dias                 ,
		lesiones_genitales                ,
		presion_arterial                  ,
		tiene_vac_antitetanica            ,
		tiene_vac_antiamerilica           ,
		tiene_vac_antihepatitis_b         ,
		tiene_encias                      ,
		tiene_caries                      ,
		tiene_edentulismo_parcial         ,
		tiene_edentulismo_total           ,
		tiene_urg_tratamiento_bucal       ,
		tiene_ansiedad                    ,
		tiene_exam_visual                 ,
		tiene_exam_colesterol             ,
		tiene_exam_glucosa                ,
		tiene_exam_mamas                  ,
		tiene_exam_prostata               ,
		tiene_exam_pelvico_pap            ,
		tiene_exam_mamografia             ,
		tiene_hab_fisica                  ,
		tiene_hab_alcohol                 ,
		tiene_hab_drogas                  ,
		tiene_planificacion_sexual        
FROM HistoriaClinica `
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = '%s'", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}

	error_find := db.QueryRow(ctx, q).Scan(
		&oHistoriaClinica.Id,
		&oHistoriaClinica.IdMedico,
		&oHistoriaClinica.IdPaciente,
		&oHistoriaClinica.EstadoCivil,
		&oHistoriaClinica.GradoInstitucion,
		&oHistoriaClinica.Ocupacion,
		&oHistoriaClinica.Direccion,
		&oHistoriaClinica.IdDistrito,
		&oHistoriaClinica.FechaRegistro,
		&oHistoriaClinica.TuvoTuberculosis,
		&oHistoriaClinica.TieneInfTransSex,
		&oHistoriaClinica.TieneDiabetes,
		&oHistoriaClinica.TieneHta,
		&oHistoriaClinica.TieneSobrepeso,
		&oHistoriaClinica.TieneInfarto,
		&oHistoriaClinica.TieneDislipenia,
		&oHistoriaClinica.TieneInfRenalGlaucoma,
		&oHistoriaClinica.TieneDepresionEsquizofrenia,
		&oHistoriaClinica.Antecedentes,
		&oHistoriaClinica.TieneHospitaliacionTransfusiones,
		&oHistoriaClinica.Dispacidad,
		&oHistoriaClinica.TieneConsumoTabaco,
		&oHistoriaClinica.TieneConsumoAlcohol,
		&oHistoriaClinica.TieneConsumoDrogas,
		&oHistoriaClinica.TieneInterQuirurjica,
		&oHistoriaClinica.Cancer,
		&oHistoriaClinica.TieneRiesgo,
		&oHistoriaClinica.TieneViolencia,
		&oHistoriaClinica.TieneSida,
		&oHistoriaClinica.TieneITS,
		&oHistoriaClinica.TieneHepatitis,
		&oHistoriaClinica.TieneDbm,
		&oHistoriaClinica.TieneCancer,
		&oHistoriaClinica.TieneDepresion,
		&oHistoriaClinica.TieneProbPsiquiatricos,
		&oHistoriaClinica.Otros,
		&oHistoriaClinica.ReaccionMedicamentos,
		&oHistoriaClinica.MedicamenteFrecuente,
		&oHistoriaClinica.EdadInicioRelacionSexual,
		&oHistoriaClinica.NumParejas,
		&oHistoriaClinica.HijosVivos,
		&oHistoriaClinica.RsMismoSexo,
		&oHistoriaClinica.Menarquia,
		&oHistoriaClinica.FlujoVagPatologico,
		&oHistoriaClinica.Dismenorrea,
		&oHistoriaClinica.TieneEmbarazo,
		&oHistoriaClinica.TieneParto,
		&oHistoriaClinica.TienePrematuro,
		&oHistoriaClinica.TieneAborto,
		&oHistoriaClinica.Gestacion,
		&oHistoriaClinica.TieneFiebre15Dias,
		&oHistoriaClinica.TieneTos15Dias,
		&oHistoriaClinica.LesionesGenitales,
		&oHistoriaClinica.PresionArterial,
		&oHistoriaClinica.TieneVacAntitetanica,
		&oHistoriaClinica.TieneVacAntiamerilica,
		&oHistoriaClinica.TieneVacAntihepatitisB,
		&oHistoriaClinica.TieneEncias,
		&oHistoriaClinica.TieneCaries,
		&oHistoriaClinica.TieneEdentulismoParcial,
		&oHistoriaClinica.TieneEdentulismoTotal,
		&oHistoriaClinica.TieneUrgTratamientoBucal,
		&oHistoriaClinica.TieneAnsiedad,
		&oHistoriaClinica.TieneExamVisual,
		&oHistoriaClinica.TieneExamColesterol,
		&oHistoriaClinica.TieneExamGlucosa,
		&oHistoriaClinica.TieneExamMamas,
		&oHistoriaClinica.TieneExamProstata,
		&oHistoriaClinica.TieneExamPelvicoPap,
		&oHistoriaClinica.TieneExamMamografia,
		&oHistoriaClinica.TieneHabFisica,
		&oHistoriaClinica.TieneHabAlcohol,
		&oHistoriaClinica.TieneHabDrogas,
		&oHistoriaClinica.TienePlanificacionSexual)

	if error_find != nil {
		return oHistoriaClinica, error_find
	}

	//Return one historia_clinica
	return oHistoriaClinica, nil
}

func Pg_FindMultiple(input_id_medico string, input_documento_identidad int, input_limit int, input_offset int) ([]models.HistoriaClinica_View, error) {

	//Initialization
	var oListHistoriaClinica []models.HistoriaClinica_View

	//Define the filters
	filters := map[string]interface{}{}
	counter_filters := 0
	if input_documento_identidad != 0 {
		filters["pa.documento_identidad"] = input_documento_identidad
		counter_filters += 1
	}
	if input_id_medico != "" {
		filters["me.id"] = input_id_medico
		counter_filters += 1
	}

	//Context timing
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//Cancel context
	defer cancel()

	//Start the connection
	db := configs.Conn_Pg_DB()

	//Define the query
	q := `SELECT hc.id,hc.fecha_Registro,pa.id,concat(pa.nombre,pa.apellido),pa.documento_identidad,me.id,concat(me.nombre,me.apellido), COALESCE(anali.id,''),to_json(array_agg(co.id::text))
	FROM historiaclinica AS hc JOIN medico AS me ON hc.id_medico=me.id JOIN paciente AS pa ON hc.id_paciente=pa.id LEFT JOIN AnalisisLaboratorio as anali ON hc.id=anali.id_historia_clinica LEFT JOIN consulta as co ON hc.id=co.id_historia_clinica`
	if counter_filters > 0 {
		q += " WHERE "
		clausulas := make([]string, 0)
		for key, value := range filters {
			clausulas = append(clausulas, fmt.Sprintf("%s = %d", key, value))
		}
		q += strings.Join(clausulas, " AND ")

	}
	rows, error_find := db.Query(ctx, q+" GROUP BY hc.id,hc.fecha_Registro,pa.id,concat(pa.nombre,pa.apellido),pa.documento_identidad,me.id,concat(me.nombre,me.apellido), anali.id ORDER BY hc.fecha_registro DESC LIMIT $1 OFFSET $2", input_limit, input_offset)
	if error_find != nil {
		return oListHistoriaClinica, error_find
	}

	//Scan the row
	for rows.Next() {
		var oHistoriaClinica models.HistoriaClinica_View
		rows.Scan(
			&oHistoriaClinica.Id,
			&oHistoriaClinica.FechaRegistro,
			&oHistoriaClinica.IdPaciente,
			&oHistoriaClinica.NombreCompletoPaciente,
			&oHistoriaClinica.DocumentoIdentidadPaciente,
			&oHistoriaClinica.IdMedico,
			&oHistoriaClinica.NombreCompletoMedico,
			&oHistoriaClinica.IdAnalisisClinico,
			&oHistoriaClinica.IdsConsultas)
		oListHistoriaClinica = append(oListHistoriaClinica, oHistoriaClinica)
	}

	if error_find != nil {
		return oListHistoriaClinica, error_find
	}

	//Return the list of establecimientos
	return oListHistoriaClinica, nil
}
