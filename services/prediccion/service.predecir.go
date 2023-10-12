package prediccion

import (
	"log"
	"strings"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	analisis_laboratorio_repository "github.com/UPC-Works/api-aliviate/repositories/analisis_laboratorio"
	historia_clinica_repository "github.com/UPC-Works/api-aliviate/repositories/historia_clinica"
	paciente_repository "github.com/UPC-Works/api-aliviate/repositories/paciente"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func Predecir(c echo.Context) error {

	//Inicilization
	var input_prediccion *models.Prediccion

	//Bind the model
	err := c.Bind(&input_prediccion)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_prediccion.IdHistoriaClinica == "" {
		return c.JSON(400, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Debe enviar el id de la historia clinica",
			},
			Data: nil})
	}

	//Get one historia_clinicas
	historia_clinica, error_find_historia_clinica := historia_clinica_repository.Pg_FindOne(input_prediccion.IdHistoriaClinica)
	if error_find_historia_clinica != nil {
		return c.JSON(500, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al buscar la historia clinica, detalles:" + error_find_historia_clinica.Error(),
			},
			Data: nil})
	}

	//Get one paciente
	paciente, error_find_paciente := paciente_repository.Pg_FindOne(historia_clinica.IdPaciente, 0)
	if error_find_paciente != nil {
		return c.JSON(500, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al buscar el paciente, detalles:" + error_find_paciente.Error(),
			},
			Data: nil})
	}

	//Get one analisis_laboratorio
	analisis_laboratorio, error_find_analisis_laboratorio := analisis_laboratorio_repository.Pg_FindOne(input_prediccion.IdHistoriaClinica)
	if error_find_analisis_laboratorio != nil {
		return c.JSON(500, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al buscar el analisis de laboratorio, detalles:" + error_find_analisis_laboratorio.Error(),
			},
			Data: nil})
	}

	// Create excel file
	file := excelize.NewFile()

	// Add an excel
	sheetName := "Sheet1"

	// Complete data
	data := map[string]interface{}{
		"ID_HISTORIA_CLINICA":                 historia_clinica.Id,
		"NOMBRES":                             paciente.Nombre,
		"APELLIDOS":                           paciente.Apellido,
		"EDAD":                                15,
		"FECHA_NACIMIENTO":                    paciente.FechaNacimiento,
		"SEXO":                                map[int]string{1: "M", 2: "F"}[paciente.Genero],
		"GRUPO_SANGUINEO":                     paciente.GrupoSanguineo,
		"GRADO_INSTITUCION":                   map[int]string{1: "Preescolar", 2: "Primaria", 3: "Secundaria", 4: "Preparatoria", 5: "Universitaria", 6: "Tecnica", 7: "Maestria", 8: "Doctorado"}[historia_clinica.GradoInstitucion],
		"ESTADO_CIVIL":                        map[int]string{1: "Soltero", 2: "Casado", 3: "Viudo", 4: "Divorciado"}[historia_clinica.EstadoCivil],
		"FECHA_REGISTRO":                      historia_clinica.FechaRegistro,
		"DIRECCION":                           historia_clinica.Direccion,
		"TUVO_TUBERCULOSIS":                   map[bool]string{true: "Si", false: "No"}[historia_clinica.TuvoTuberculosis],
		"TIENE_INF_RENAL_GLAUCOMA":            map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfRenalGlaucoma],
		"TIENE_INF_TRANS_SEX":                 map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfTransSex],
		"TIENE_SIDA":                          map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneSida],
		"TIENE_ITS":                           map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneITS],
		"TIENE_HEPATITIS":                     map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHepatitis],
		"TIENE_DIABETES":                      map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDiabetes],
		"TIENE_HTA":                           map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHta],
		"TIENE_SOBREPESO":                     map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneSobrepeso],
		"TIENE_DISLIPENIA":                    map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDislipenia],
		"MENARQUIA":                           "No",
		"TIENE_DEPRESION_ESQUIZOFRENIA":       map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDepresionEsquizofrenia],
		"TIENE_HOSPITALIZACION_TRANSFUCIONES": map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHospitaliacionTransfusiones],
		"TIENE_INTER_QUIRURJICA":              map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInterQuirurjica],
		"TIENE_PREMATURO":                     map[bool]string{true: "Si", false: "No"}[historia_clinica.TienePrematuro],
		"TIENE_ABORTO":                        map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneAborto],
		"TIENE_PARTO":                         map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneParto],
		"FLUJO_VAG_PATOLOGICO":                map[bool]string{true: "Si", false: "No"}[historia_clinica.FlujoVagPatologico],
		"GESTACION":                           "",
		"TIENE_EXAM_PROSTATA":                 map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamProstata],
		"TIENE_VIOLENCIA":                     map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneViolencia],
		"TIENE_DBM":                           map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDbm],
		"TIENE_INFARTO":                       map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfarto],
		"TIENE_CANCER":                        map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneCancer],
		"TIENE_DEPRESION":                     map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDepresion],
		"TIENE_PROB_PSIQUIATRICOS":            map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneProbPsiquiatricos],
		"RS_MISMO_SEXO":                       map[bool]string{true: "Si", false: "No"}[historia_clinica.RsMismoSexo],
		"MEDICAMENTO_FRECUENTE":               strings.Join(historia_clinica.MedicamenteFrecuente, ","),
		"REACCION_MEDICAMENTOS":               strings.Join(historia_clinica.ReaccionMedicamentos, ","),
		"TIENE_CONSUMO_TABACO":                map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneConsumoTabaco],
		"TIENE_CONSUMO_ALCOHOL":               map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneConsumoAlcohol],
		"EDAD_INI_RELACION_SEXUAL":            historia_clinica.EdadInicioRelacionSexual,
		"NUM_PAREJAS":                         historia_clinica.NumParejas,
		"FECHA_ULT_REGLA":                     "",
		"DISMENORREA":                         map[bool]string{true: "Si", false: "No"}[historia_clinica.Dismenorrea],
		"TIENE_EMBARAZO":                      map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEmbarazo],
		"PRESION_ARTERIAL":                    historia_clinica.PresionArterial,
		"LESIONES_GENITALES":                  "No",
		"TIENE_FIEBRE_15_DIAS":                map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneFiebre15Dias],
		"TIENE_TOS_15_DIAS":                   map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneTos15Dias],
		"TIENE_VAC_ANTITETANICA":              map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntitetanica],
		"TIENE_VAC_ANTIAMERILICA":             map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntiamerilica],
		"TIENE_VAC_ANTIHEPATITIS_B":           map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntihepatitisB],
		"TIENE_ENCIAS":                        map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEncias],
		"TIENE_CARIES":                        map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneCaries],
		"TIENE_EDENTULISMO_TOTAL":             map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEdentulismoTotal],
		"TIENE_ANSIEDAD":                      map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneAnsiedad],
		"TIENE_EDENTULISMO_PARCIAL":           map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEdentulismoParcial],
		"TIENE_EXAM_VISUAL":                   map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamVisual],
		"TIENE_URG_TRATAMIENTO_BUCAL":         map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneUrgTratamientoBucal],
		"TIENE_MAMOGRAFIA":                    map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamMamografia],
		"TIENE_EXAM_PELVICO_PAP":              map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamPelvicoPap],
		"TIENE_EXAM_COLESTEROL":               map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamColesterol],
		"TIENE_EXAM_MAMAS":                    map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamMamas],
		"TIENE_EXAM_GLUCOSA":                  map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamGlucosa],
		"TIENE_HAB_FISICA":                    map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabFisica],
		"TIENE_PLANIFICACION_SEXUAL":          map[bool]string{true: "Si", false: "No"}[historia_clinica.TienePlanificacionSexual],
		"TIENE_HAB_ALCOHOL":                   map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabAlcohol],
		"TIENE_HAB_DROGAS":                    map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabDrogas],
		"HEMATOCRITO":                         analisis_laboratorio.Hematrocito,
		"HEMOGLOBINA":                         analisis_laboratorio.Hemoglobina,
		"COLESTEROL":                          analisis_laboratorio.Colesterol,
		"TRIGLICERIDOS":                       analisis_laboratorio.Trigliceridos,
		"COLESTEROL_HDL":                      analisis_laboratorio.ColesterolHdl,
		"COLESTEROL_LDL":                      analisis_laboratorio.ColesterolLdl,
		"COLESTEROL_VLDL":                     analisis_laboratorio.ColesterolVldl,
		"RIESGO_1":                            analisis_laboratorio.Riesgo1,
		"RIESGO_2":                            analisis_laboratorio.Riesgo2,
		"GLUCOSA":                             analisis_laboratorio.Glucosa,
		"P_A":                                 input_prediccion.ConsultaActual.PA,
		"F_C":                                 input_prediccion.ConsultaActual.FC,
		"F_R":                                 input_prediccion.ConsultaActual.FR,
		"I_M_C":                               input_prediccion.ConsultaActual.IMC,
		"TEMPERATURA":                         input_prediccion.ConsultaActual.Temperatura,
		"PESO":                                input_prediccion.ConsultaActual.Peso,
		"TALLA":                               input_prediccion.ConsultaActual.Talla,
		"ENFERMEDAD":                          "",
		"SIGNOS_SINTOMAS":                     input_prediccion.ConsultaActual.SignosSintomas,
	}

	// Fill the excel
	for cell, value := range data {
		file.SetCellValue(sheetName, cell, value)
	}

	// Guardar el archivo
	if err := file.SaveAs("ejemplo.xlsx"); err != nil {
		log.Println("Error: ", err)
	} else {
		log.Println("Archivo Excel creado con éxito.")
	}

	//OK
	return c.JSON(200, &helpers.ResponseString{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: "OK"})
}
