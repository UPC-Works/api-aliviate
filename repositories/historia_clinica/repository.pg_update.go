package historia_clinica

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Update(input_historia_clinica *models.HistoriaClinica) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `UPDATE HistoriaClinica SET
		id_medico=$1,
		id_paciente=$2,
		estado_civil=$3,
		grado_institucion=$4,
		ocupacion=$5,
		direccion=$6,
		id_distrito=$7                       ,
		fecha_registro=$8                    ,
		tuvo_tuberculosis=$9                 ,
		tiene_inf_trans_sex=$10               ,
		tiene_diabetes=$11                    ,
		tiene_hta=$12                          ,
		tiene_sobrepeso=$13                   ,
		tiene_infarto=$14                      ,
		tiene_dislipenia=$15                   ,
		tiene_inf_renal_glaucoma=$16           ,
		tiene_depresion_esquizofrenia=$17      ,
		antecedentes=$18                       ,
		tiene_hospitaliacion_transfusiones=$19 ,
		dispacidad=$20                         ,
		tiene_consumo_tabaco=$21               ,
		tiene_consumo_alcohol=$22              ,
		tiene_consumo_drogas=$23               ,
		tiene_inter_quirurjica=$24             ,
		cancer=$25                             ,
		tiene_riesgo=$26                       ,
		tiene_violencia=$27                    ,
		tiene_sida=$28                         ,
		tiene_its=$29                          ,
		tiene_hepatitis=$30                    ,
		tiene_dbm=$31                          ,
		tiene_cancer=$32                      ,
		tiene_depresion=$33                    ,
		tiene_prob_psiquiatricos=$34           ,
		otros=$35                              ,
		reaccion_medicamentos=$36              ,
		medicamente_frecuente=$37              ,
		edad_inicio_relacion_sexual=$38        ,
		num_parejas=$39                        ,
		hijos_vivos=$40                        ,
		rs_mismo_sexo=$41                      ,
		menarquia=$42                          ,
		flujo_vag_patologico=$43               ,
		dismenorrea=$44                        ,
		tiene_embarazo=$45                     ,
		tiene_parto=$46                        ,
		tiene_prematuro=$47                    ,
		tiene_aborto=$48                       ,
		gestacion=$49                          ,
		tiene_fiebre_15_dias=$50               ,
		tiene_tos_15_dias=$51                  ,
		lesiones_genitales=$52                 ,
		presion_arterial=$53                   ,
		tiene_vac_antitetanica=$54             ,
		tiene_vac_antiamerilica=$55            ,
		tiene_vac_antihepatitis_b=$56          ,
		tiene_encias=$57                       ,
		tiene_caries=$58                      ,
		tiene_edentulismo_parcial=$59          ,
		tiene_edentulismo_total=$60            ,
		tiene_urg_tratamiento_bucal=$61        ,
		tiene_ansiedad=$62                    ,
		tiene_exam_visual=$63                  ,
		tiene_exam_colesterol=$64              ,
		tiene_exam_glucosa=$65                 ,
		tiene_exam_mamas=$66                   ,
		tiene_exam_prostata=$67                ,
		tiene_exam_pelvico_pap=$68             ,
		tiene_exam_mamografia=$69              ,
		tiene_hab_fisica=$70                   ,
		tiene_hab_alcohol=$71                  ,
		tiene_hab_drogas=$72                   ,
		tiene_planificacion_sexual=$73         
	 WHERE id                                =$74`
	_, err_query := db.Exec(ctx, query,
		input_historia_clinica.IdMedico,
		input_historia_clinica.IdPaciente,
		input_historia_clinica.EstadoCivil,
		input_historia_clinica.GradoInstitucion,
		input_historia_clinica.Ocupacion,
		input_historia_clinica.Direccion,
		input_historia_clinica.IdDistrito,
		input_historia_clinica.FechaRegistro,
		input_historia_clinica.TuvoTuberculosis,
		input_historia_clinica.TieneInfTransSex,
		input_historia_clinica.TieneDiabetes,
		input_historia_clinica.TieneHta,
		input_historia_clinica.TieneSobrepeso,
		input_historia_clinica.TieneInfarto,
		input_historia_clinica.TieneDislipenia,
		input_historia_clinica.TieneInfRenalGlaucoma,
		input_historia_clinica.TieneDepresionEsquizofrenia,
		input_historia_clinica.Antecedentes,
		input_historia_clinica.TieneHospitaliacionTransfusiones,
		input_historia_clinica.Dispacidad,
		input_historia_clinica.TieneConsumoTabaco,
		input_historia_clinica.TieneConsumoAlcohol,
		input_historia_clinica.TieneConsumoDrogas,
		input_historia_clinica.TieneInterQuirurjica,
		input_historia_clinica.Cancer,
		input_historia_clinica.TieneRiesgo,
		input_historia_clinica.TieneViolencia,
		input_historia_clinica.TieneSida,
		input_historia_clinica.TieneITS,
		input_historia_clinica.TieneHepatitis,
		input_historia_clinica.TieneDbm,
		input_historia_clinica.TieneCancer,
		input_historia_clinica.TieneDepresion,
		input_historia_clinica.TieneProbPsiquiatricos,
		input_historia_clinica.Otros,
		input_historia_clinica.ReaccionMedicamentos,
		input_historia_clinica.MedicamenteFrecuente,
		input_historia_clinica.EdadInicioRelacionSexual,
		input_historia_clinica.NumParejas,
		input_historia_clinica.HijosVivos,
		input_historia_clinica.RsMismoSexo,
		input_historia_clinica.Menarquia,
		input_historia_clinica.FlujoVagPatologico,
		input_historia_clinica.Dismenorrea,
		input_historia_clinica.TieneEmbarazo,
		input_historia_clinica.TieneParto,
		input_historia_clinica.TienePrematuro,
		input_historia_clinica.TieneAborto,
		input_historia_clinica.Gestacion,
		input_historia_clinica.TieneFiebre15Dias,
		input_historia_clinica.TieneTos15Dias,
		input_historia_clinica.LesionesGenitales,
		input_historia_clinica.PresionArterial,
		input_historia_clinica.TieneVacAntitetanica,
		input_historia_clinica.TieneVacAntiamerilica,
		input_historia_clinica.TieneVacAntihepatitisB,
		input_historia_clinica.TieneEncias,
		input_historia_clinica.TieneCaries,
		input_historia_clinica.TieneEdentulismoParcial,
		input_historia_clinica.TieneEdentulismoTotal,
		input_historia_clinica.TieneUrgTratamientoBucal,
		input_historia_clinica.TieneAnsiedad,
		input_historia_clinica.TieneExamVisual,
		input_historia_clinica.TieneExamColesterol,
		input_historia_clinica.TieneExamGlucosa,
		input_historia_clinica.TieneExamMamas,
		input_historia_clinica.TieneExamProstata,
		input_historia_clinica.TieneExamPelvicoPap,
		input_historia_clinica.TieneExamMamografia,
		input_historia_clinica.TieneHabFisica,
		input_historia_clinica.TieneHabAlcohol,
		input_historia_clinica.TieneHabDrogas,
		input_historia_clinica.TienePlanificacionSexual,
		input_historia_clinica.Id,
	)

	if err_query != nil {
		return err_query
	}

	return nil
}