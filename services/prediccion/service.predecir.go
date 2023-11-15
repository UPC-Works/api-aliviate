package prediccion

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tealeg/xlsx"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	historia_clinica_repository "github.com/UPC-Works/api-aliviate/repositories/historia_clinica"
	paciente_repository "github.com/UPC-Works/api-aliviate/repositories/paciente"
)

func Predecir(c echo.Context) error {

	//Inicilization
	var input_consulta *models.Consulta

	//Bind the model
	err := c.Bind(&input_consulta)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Revisa la estructura o el tipo del valor",
			},
			Data: ""})
	}

	//Validation of the Business Rules
	if input_consulta.IdHistoriaClinica == "" {
		return c.JSON(400, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Debe enviar el id de la historia clinica",
			},
			Data: nil})
	}

	//Get one historia_clinicas
	historia_clinica, error_find_historia_clinica := historia_clinica_repository.Pg_FindOne(input_consulta.IdHistoriaClinica)
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
	/*analisis_laboratorio, error_find_analisis_laboratorio := analisis_laboratorio_repository.Pg_FindOne(input_consulta.IdHistoriaClinica)
	if error_find_analisis_laboratorio != nil {
		return c.JSON(500, &helpers.ResponsePrediccion{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Error al buscar el analisis de laboratorio, detalles:" + error_find_analisis_laboratorio.Error(),
			},
			Data: nil})
	}*/

	// Crear un nuevo archivo Excel
	file := xlsx.NewFile()

	// Agregar una hoja al archivo
	sheet, err := file.AddSheet("Hoja1")
	if err != nil {
		panic(err)
	}

	// Crear una fila en la hoja
	row := sheet.AddRow()

	// Agregar celdas a la fila
	cell := row.AddCell()
	cell.Value = "ID_HISTORIA_CLINICA"

	cell = row.AddCell()
	cell.Value = "NOMBRES"

	cell = row.AddCell()
	cell.Value = "APELLIDOS"
	cell = row.AddCell()
	cell.Value = "EDAD"
	cell = row.AddCell()
	cell.Value = "FECHA_NACIMIENTO"
	cell = row.AddCell()
	cell.Value = "SEXO"
	cell = row.AddCell()
	cell.Value = "GRUPO_SANGUINEO"
	cell = row.AddCell()
	cell.Value = "GRADO_INSTITUCION"
	cell = row.AddCell()
	cell.Value = "ESTADO_CIVIL"
	cell = row.AddCell()
	cell.Value = "FECHA_REGISTRO"
	cell = row.AddCell()
	cell.Value = "DIRECCION"
	cell = row.AddCell()
	cell.Value = "TUVO_TUBERCULOSIS"
	cell = row.AddCell()
	cell.Value = "TIENE_INF_RENAL_GLAUCOMA"
	cell = row.AddCell()
	cell.Value = "TIENE_INF_TRANS_SEX"
	cell = row.AddCell()
	cell.Value = "TIENE_SIDA"
	cell = row.AddCell()
	cell.Value = "TIENE_ITS"
	cell = row.AddCell()
	cell.Value = "TIENE_HEPATITIS"
	cell = row.AddCell()
	cell.Value = "TIENE_DIABETES"
	cell = row.AddCell()
	cell.Value = "TIENE_HTA"
	cell = row.AddCell()
	cell.Value = "TIENE_SOBREPESO"
	cell = row.AddCell()
	cell.Value = "TIENE_DISLIPENIA"
	cell = row.AddCell()
	cell.Value = "MENARQUIA"
	cell = row.AddCell()
	cell.Value = "TIENE_DEPRESION_ESQUIZOFRENIA"
	cell = row.AddCell()
	cell.Value = "TIENE_HOSPITALIZACION_TRANSFUCIONES"
	cell = row.AddCell()
	cell.Value = "TIENE_INTER_QUIRURJICA"
	cell = row.AddCell()
	cell.Value = "TIENE_PREMATURO"
	cell = row.AddCell()
	cell.Value = "TIENE_ABORTO"
	cell = row.AddCell()
	cell.Value = "TIENE_PARTO"
	cell = row.AddCell()
	cell.Value = "FLUJO_VAG_PATOLOGICO"
	cell = row.AddCell()
	cell.Value = "GESTACION"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_PROSTATA"
	cell = row.AddCell()
	cell.Value = "TIENE_VIOLENCIA"
	cell = row.AddCell()
	cell.Value = "TIENE_DBM"
	cell = row.AddCell()
	cell.Value = "TIENE_INFARTO"
	cell = row.AddCell()
	cell.Value = "TIENE_CANCER"
	cell = row.AddCell()
	cell.Value = "TIENE_DEPRESION"
	cell = row.AddCell()
	cell.Value = "TIENE_PROB_PSIQUIATRICOS"
	cell = row.AddCell()
	cell.Value = "RS_MISMO_SEXO"
	cell = row.AddCell()
	cell.Value = "MEDICAMENTO_FRECUENTE"
	cell = row.AddCell()
	cell.Value = "REACCION_MEDICAMENTOS"
	cell = row.AddCell()
	cell.Value = "TIENE_CONSUMO_TABACO"
	cell = row.AddCell()
	cell.Value = "TIENE_CONSUMO_ALCOHOL"
	cell = row.AddCell()
	cell.Value = "EDAD_INI_RELACION_SEXUAL"
	cell = row.AddCell()
	cell.Value = "NUM_PAREJAS"
	cell = row.AddCell()
	cell.Value = "FECHA_ULT_REGLA"
	cell = row.AddCell()
	cell.Value = "DISMENORREA"
	cell = row.AddCell()
	cell.Value = "TIENE_EMBARAZO"
	cell = row.AddCell()
	cell.Value = "PRESION_ARTERIAL"
	cell = row.AddCell()
	cell.Value = "LESIONES_GENITALES"
	cell = row.AddCell()
	cell.Value = "TIENE_FIEBRE_15_DIAS"
	cell = row.AddCell()
	cell.Value = "TIENE_TOS_15_DIAS"
	cell = row.AddCell()
	cell.Value = "TIENE_VAC_ANTITETANICA"
	cell = row.AddCell()
	cell.Value = "TIENE_VAC_ANTIAMERILICA"
	cell = row.AddCell()
	cell.Value = "TIENE_VAC_ANTIHEPATITIS_B"
	cell = row.AddCell()
	cell.Value = "TIENE_ENCIAS"
	cell = row.AddCell()
	cell.Value = "TIENE_CARIES"
	cell = row.AddCell()
	cell.Value = "TIENE_EDENTULISMO_TOTAL"
	cell = row.AddCell()
	cell.Value = "TIENE_ANSIEDAD"
	cell = row.AddCell()
	cell.Value = "TIENE_EDENTULISMO_PARCIAL"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_VISUAL"
	cell = row.AddCell()
	cell.Value = "TIENE_URG_TRATAMIENTO_BUCAL"
	cell = row.AddCell()
	cell.Value = "TIENE_MAMOGRAFIA"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_PELVICO_PAP"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_COLESTEROL"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_MAMAS"
	cell = row.AddCell()
	cell.Value = "TIENE_EXAM_GLUCOSA"
	cell = row.AddCell()
	cell.Value = "TIENE_HAB_FISICA"
	cell = row.AddCell()
	cell.Value = "TIENE_PLANIFICACION_SEXUAL"
	cell = row.AddCell()
	cell.Value = "TIENE_HAB_ALCOHOL"
	cell = row.AddCell()
	cell.Value = "TIENE_HAB_DROGAS"
	cell = row.AddCell()
	cell.Value = "HEMATOCRITO"
	cell = row.AddCell()
	cell.Value = "HEMOGLOBINA"
	cell = row.AddCell()
	cell.Value = "COLESTEROL"
	cell = row.AddCell()
	cell.Value = "TRIGLICERIDOS"
	cell = row.AddCell()
	cell.Value = "COLESTEROL_HDL"
	cell = row.AddCell()
	cell.Value = "COLESTEROL_LDL"
	cell = row.AddCell()
	cell.Value = "COLESTEROL_VLDL"
	cell = row.AddCell()
	cell.Value = "RIESGO_1"
	cell = row.AddCell()
	cell.Value = "RIESGO_2"
	cell = row.AddCell()
	cell.Value = "GLUCOSA"
	cell = row.AddCell()
	cell.Value = "P_A"
	cell = row.AddCell()
	cell.Value = "F_C"
	cell = row.AddCell()
	cell.Value = "F_R"
	cell = row.AddCell()
	cell.Value = "I_M_C"
	cell = row.AddCell()
	cell.Value = "TEMPERATURA"
	cell = row.AddCell()
	cell.Value = "PESO"
	cell = row.AddCell()
	cell.Value = "TALLA"
	cell = row.AddCell()
	cell.Value = "ENFERMEDAD"
	cell = row.AddCell()
	cell.Value = "SIGNOS_SINTOMAS"

	// Agregar celdas a la fila
	dataRow := sheet.AddRow()
	cell = dataRow.AddCell()
	cell.Value = historia_clinica.Id

	cell = dataRow.AddCell()
	cell.Value = paciente.Nombre

	cell = dataRow.AddCell()
	cell.Value = paciente.Apellido
	cell = dataRow.AddCell()
	cell.SetInt(15)
	cell = dataRow.AddCell()
	cell.Value = "O"
	cell = dataRow.AddCell()
	cell.Value = map[int]string{1: "M", 2: "F"}[paciente.Genero]
	cell = dataRow.AddCell()
	cell.Value = paciente.GrupoSanguineo
	cell = dataRow.AddCell()
	cell.Value = map[int]string{1: "Preescolar", 2: "Primaria", 3: "Secundaria", 4: "Preparatoria", 5: "Universitaria", 6: "Tecnica", 7: "Maestria", 8: "Doctorado"}[historia_clinica.GradoInstitucion]
	cell = dataRow.AddCell()
	cell.Value = map[int]string{1: "Soltero", 2: "Casado", 3: "Viudo", 4: "Divorciado"}[historia_clinica.EstadoCivil]
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = historia_clinica.Direccion
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TuvoTuberculosis]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfRenalGlaucoma]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfTransSex]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneSida]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneITS]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHepatitis]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDiabetes]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHta]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneSobrepeso]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDislipenia]
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDepresionEsquizofrenia]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHospitaliacionTransfusiones]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInterQuirurjica]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TienePrematuro]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneAborto]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneParto]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.FlujoVagPatologico]
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamProstata]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneViolencia]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDbm]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneInfarto]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneCancer]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneDepresion]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneProbPsiquiatricos]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.RsMismoSexo]
	cell = dataRow.AddCell()
	cell.Value = strings.Join(historia_clinica.MedicamenteFrecuente, ",")
	cell = dataRow.AddCell()
	cell.Value = strings.Join(historia_clinica.ReaccionMedicamentos, ",")
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneConsumoTabaco]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneConsumoAlcohol]
	cell = dataRow.AddCell()
	cell.SetInt(historia_clinica.EdadInicioRelacionSexual)
	cell = dataRow.AddCell()
	cell.SetInt(historia_clinica.NumParejas)
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.Dismenorrea]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEmbarazo]
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = "No"
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneFiebre15Dias]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneTos15Dias]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntitetanica]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntiamerilica]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneVacAntihepatitisB]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEncias]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneCaries]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEdentulismoTotal]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneAnsiedad]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneEdentulismoParcial]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamVisual]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneUrgTratamientoBucal]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamMamografia]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamPelvicoPap]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamColesterol]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamMamas]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneExamGlucosa]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabFisica]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TienePlanificacionSexual]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabAlcohol]
	cell = dataRow.AddCell()
	cell.Value = map[bool]string{true: "Si", false: "No"}[historia_clinica.TieneHabDrogas]
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Hematrocito))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Hemoglobina))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Colesterol))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Trigliceridos))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.ColesterolHdl))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.ColesterolLdl))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.ColesterolVldl))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Riesgo1))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Riesgo2))
	cell = dataRow.AddCell()
	//cell.SetFloat(float64(analisis_laboratorio.Glucosa))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.PA))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.FC))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.FR))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.IMC))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.Temperatura))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.Peso))
	cell = dataRow.AddCell()
	cell.SetFloat(float64(input_consulta.Talla))
	cell = dataRow.AddCell()
	cell.Value = "Na"
	cell = dataRow.AddCell()
	cell.Value = input_consulta.SignosSintomas

	// Guardar el archivo Excel
	nombre_archivo := uuid.New().String()
	err = file.Save(nombre_archivo + ".xlsx")
	if err != nil {
		panic(err)
	}

	//Enviar a predecir
	//respuesta_prediccion := EnviarArchivo(nombre_archivo)

	//Eliminar archivo
	err_eliminar := os.Remove(nombre_archivo + ".xlsx")
	if err_eliminar != nil {
		fmt.Println("Error al eliminar el archivo:", err_eliminar)
	} else {
		fmt.Println("Archivo eliminado exitosamente.")
	}

	predicciones := []models.PrediccionEnfermedadShow{
		{
			Enfermedad:   "Resfriado",
			Probabilidad: 0.75,
		},
		{
			Enfermedad:   "Gripe",
			Probabilidad: 0.95,
		},
		{
			Enfermedad:   "Dolor de garganta",
			Probabilidad: 0.60,
		},
	}

	//OK
	return c.JSON(200, &helpers.ResponsePrediccion{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: predicciones})
}
