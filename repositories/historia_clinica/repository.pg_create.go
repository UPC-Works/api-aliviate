package historia_clinica

import (
	"context"
	"time"

	configs "github.com/UPC-Works/api-aliviate/configs"
	models "github.com/UPC-Works/api-aliviate/models"
)

func Pg_Create(input_historia_clinica *models.HistoriaClinica) error {

	//Context time limit
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	db := configs.Conn_Pg_DB()

	query := `INSERT INTO Medico (
		id                                ,
		id_paciente,
		id_distrito                       ,
		fecha_registro                    ,
		nombres                           ,
		apellidos                         ,
		fecha_nacimiento                  ,
		genero                            ,
		grupo_sanguineo                   ,
		rh_sanguineo                      ,
		grado_institucion                 ,
		estado_civil                      ,
		ocupación                         ,
		documento_identidad               ,
		direccion                         ,
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
		tiene_infarto                     ,
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
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29,$30,$31,$32,$33,$34,$35,$36,$37,$38,$39,,$40,$41,$42,$43,$44,$45,$46,$47,$48,$49,$50,$51,$52,$53,$54,$55,$56,$57,$58,$59,$60,$61,$62,$63,$64,$65,$66,$67,$68,$69,$70,$71,$72,$73,$74)`
	_, err_query := db.Exec(ctx, query,
		input_historia_clinica.Id,
		input_historia_clinica.IdPaciente,
		input_historia_clinica.IdDistrito,
		input_historia_clinica.FechaRegistro,
		input_historia_clinica.Direccion,
		input_historia_clinica.GrupoSanguineo,
		input_historia_clinica.RhSanguineo,
		input_historia_clinica.GradoInstitucion,
		input_historia_clinica.RstadoCivil,
		input_historia_clinica.Ocupación,
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
	)

	if err_query != nil {
		return err_query
	}

	return nil
}
