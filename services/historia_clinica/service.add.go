package historia_clinica

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	helpers "github.com/UPC-Works/api-aliviate/helpers"
	models "github.com/UPC-Works/api-aliviate/models"
	historia_clinica_repository "github.com/UPC-Works/api-aliviate/repositories/historia_clinica"
)

func Add(c echo.Context) error {

	//Inicilization
	var input_historia_clinica *models.HistoriaClinica

	//Bind the model
	err := c.Bind(&input_historia_clinica)
	if err != nil {
		return c.JSON(400, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   "Check the structure or the type of the value",
			},
			Data: ""})
	}

	//Storage the Historia Clinica
	new_historia_clinica := models.NewHistoriaClinica(uuid.New().String(), input_historia_clinica.IdPaciente, input_historia_clinica.IdDistrito, input_historia_clinica.FechaRegistro, input_historia_clinica.EstadoCivil, input_historia_clinica.GradoInstitucion, input_historia_clinica.Ocupacion, input_historia_clinica.Direccion, input_historia_clinica.TuvoTuberculosis, input_historia_clinica.TieneInfTransSex, input_historia_clinica.TieneDiabetes, input_historia_clinica.TieneHta, input_historia_clinica.TieneSobrepeso, input_historia_clinica.TieneInfarto, input_historia_clinica.TieneDislipenia, input_historia_clinica.TieneInfRenalGlaucoma, input_historia_clinica.TieneDepresionEsquizofrenia, input_historia_clinica.Antecedentes, input_historia_clinica.TieneHospitaliacionTransfusiones, input_historia_clinica.Dispacidad, input_historia_clinica.TieneConsumoTabaco, input_historia_clinica.TieneConsumoAlcohol, input_historia_clinica.TieneConsumoDrogas, input_historia_clinica.TieneInterQuirurjica, input_historia_clinica.Cancer, input_historia_clinica.TieneRiesgo, input_historia_clinica.TieneViolencia, input_historia_clinica.TieneSida, input_historia_clinica.TieneITS, input_historia_clinica.TieneHepatitis, input_historia_clinica.TieneDbm, input_historia_clinica.TieneCancer, input_historia_clinica.TieneDepresion, input_historia_clinica.TieneProbPsiquiatricos, input_historia_clinica.Otros, input_historia_clinica.ReaccionMedicamentos, input_historia_clinica.MedicamenteFrecuente, input_historia_clinica.EdadInicioRelacionSexual, input_historia_clinica.NumParejas, input_historia_clinica.HijosVivos, input_historia_clinica.RsMismoSexo, input_historia_clinica.Menarquia, input_historia_clinica.FlujoVagPatologico, input_historia_clinica.Dismenorrea, input_historia_clinica.TieneEmbarazo, input_historia_clinica.TieneParto, input_historia_clinica.TienePrematuro, input_historia_clinica.TieneAborto, input_historia_clinica.Gestacion, input_historia_clinica.TieneFiebre15Dias, input_historia_clinica.TieneTos15Dias, input_historia_clinica.LesionesGenitales, input_historia_clinica.PresionArterial, input_historia_clinica.TieneVacAntitetanica, input_historia_clinica.TieneVacAntiamerilica, input_historia_clinica.TieneVacAntihepatitisB, input_historia_clinica.TieneEncias, input_historia_clinica.TieneCaries, input_historia_clinica.TieneEdentulismoParcial, input_historia_clinica.TieneEdentulismoTotal, input_historia_clinica.TieneUrgTratamientoBucal, input_historia_clinica.TieneAnsiedad, input_historia_clinica.TieneExamVisual, input_historia_clinica.TieneExamColesterol, input_historia_clinica.TieneExamGlucosa, input_historia_clinica.TieneExamMamas, input_historia_clinica.TieneExamProstata, input_historia_clinica.TieneExamPelvicoPap, input_historia_clinica.TieneExamMamografia, input_historia_clinica.TieneHabFisica, input_historia_clinica.TieneHabAlcohol, input_historia_clinica.TieneHabDrogas, input_historia_clinica.TienePlanificacionSexual)
	error_create_paciente := historia_clinica_repository.Pg_Create(new_historia_clinica)
	if error_create_paciente != nil {
		return c.JSON(500, &helpers.ResponseString{
			Error: helpers.ErrorStructure{
				HasError: true,
				Detail:   error_create_paciente.Error(),
			},
			Data: ""})
	}

	//OK
	return c.JSON(200, &helpers.ResponseString{
		Error: helpers.ErrorStructure{
			HasError: false,
			Detail:   "",
		},
		Data: "OK"})
}
