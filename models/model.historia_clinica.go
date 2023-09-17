package models

import "time"

// Model

type HistoriaClinica struct {
	Id                               string      `json:"id "`
	IdMedico                         string      `json:"idMedico"`
	IdPaciente                       string      `json:"idPaciente"`
	EstadoCivil                      int         `json:"estadoCivil"`
	GradoInstitucion                 int         `json:"gradoInstitucion"`
	Ocupacion                        string      `json:"ocupacion"`
	Direccion                        string      `json:"direccion"`
	IdDistrito                       int         `json:"idDistrito"`
	FechaRegistro                    time.Time   `json:"fechaRegistro"`
	TuvoTuberculosis                 bool        `json:"tuvoTuberculosis"`
	TieneInfTransSex                 bool        `json:"tieneInfTransSex"`
	TieneDiabetes                    bool        `json:"tieneDiabetes"`
	TieneHta                         bool        `json:"tieneHta"`
	TieneSobrepeso                   bool        `json:"tieneSobrepeso"`
	TieneInfarto                     bool        `json:"tieneInfarto"`
	TieneDislipenia                  bool        `json:"tieneDislipenia"`
	TieneInfRenalGlaucoma            bool        `json:"tieneInfRenalGlaucoma"`
	TieneDepresionEsquizofrenia      bool        `json:"tieneDepresionEsquizofrenia"`
	Antecedentes                     []string    `json:"antecedentes"`
	TieneHospitaliacionTransfusiones bool        `json:"tieneHospitaliacionTransfusiones"`
	Dispacidad                       []string    `json:"dispacidad"`
	TieneConsumoTabaco               bool        `json:"tieneConsumoTabaco"`
	TieneConsumoAlcohol              bool        `json:"tieneConsumoAlcohol"`
	TieneConsumoDrogas               bool        `json:"tieneConsumoDrogas"`
	TieneInterQuirurjica             bool        `json:"tieneInterQuirurjica"`
	Cancer                           []string    `json:"cancer"`
	TieneRiesgo                      bool        `json:"tieneRiesgo"`
	TieneViolencia                   bool        `json:"tieneViolencia"`
	TieneSida                        bool        `json:"tieneSid"`
	TieneITS                         bool        `json:"tieneITS"`
	TieneHepatitis                   bool        `json:"tieneHepatitis"`
	TieneDbm                         bool        `json:"tieneDbm"`
	TieneCancer                      bool        `json:"tieneCancer"`
	TieneDepresion                   bool        `json:"tieneDepresion"`
	TieneProbPsiquiatricos           bool        `json:"tieneProbPsiquiatricos"`
	Otros                            []string    `json:"otros"`
	ReaccionMedicamentos             []string    `json:"reaccionMedicamentos"`
	MedicamenteFrecuente             []string    `json:"medicamenteFrecuente"`
	EdadInicioRelacionSexual         int         `json:"edadInicioRelacionSexual"`
	NumParejas                       int         `json:"numParejas"`
	HijosVivos                       int         `json:"hijosVivos"`
	RsMismoSexo                      bool        `json:"rsMismoSexo"`
	Menarquia                        interface{} `json:"menarquia"`
	FlujoVagPatologico               bool        `json:"flujoVagPatologico"`
	Dismenorrea                      bool        `json:"dismenorrea"`
	TieneEmbarazo                    bool        `json:"tieneEmbarazo"`
	TieneParto                       bool        `json:"tieneParto"`
	TienePrematuro                   bool        `json:"tienePrematuro"`
	TieneAborto                      bool        `json:"tieneAborto"`
	Gestacion                        interface{} `json:"gestacion"`
	TieneFiebre15Dias                bool        `json:"tieneFiebre15Dias"`
	TieneTos15Dias                   bool        `json:"tieneTos15Dias"`
	LesionesGenitales                []string    `json:"lesionesGenitales"`
	PresionArterial                  interface{} `json:"presionArterial"`
	TieneVacAntitetanica             bool        `json:"tieneVacAntitetanica "`
	TieneVacAntiamerilica            bool        `json:"tieneVacAntiamerilica"`
	TieneVacAntihepatitisB           bool        `json:"tieneVacAntihepatitisB"`
	TieneEncias                      bool        `json:"tieneEncias"`
	TieneCaries                      bool        `json:"tieneCaries"`
	TieneEdentulismoParcial          bool        `json:"tieneEdentulismoParcial"`
	TieneEdentulismoTotal            bool        `json:"tieneEdentulismoTotal"`
	TieneUrgTratamientoBucal         bool        `json:"tieneUrgTratamientoBucal"`
	TieneAnsiedad                    bool        `json:"tieneAnsiedad"`
	TieneExamVisual                  bool        `json:"tieneExamVisual"`
	TieneExamColesterol              bool        `json:"tieneExamColesterol"`
	TieneExamGlucosa                 bool        `json:"tieneExamGlucosa"`
	TieneExamMamas                   bool        `json:"tieneExamMamas"`
	TieneExamProstata                bool        `json:"tieneExamProstata"`
	TieneExamPelvicoPap              bool        `json:"tieneExamPelvicoPap"`
	TieneExamMamografia              bool        `json:"tieneExamMamografia"`
	TieneHabFisica                   bool        `json:"tieneHabFisica"`
	TieneHabAlcohol                  bool        `json:"tieneHabAlcohol"`
	TieneHabDrogas                   bool        `json:"tieneHabDrogas"`
	TienePlanificacionSexual         bool        `json:"tienePlanificacionSexual"`
}

//Constructor

func NewHistoriaClinica(
	id string, idMedico string, idPaciente string, idDistrito int, fechaRegistro time.Time, estadoCivil int, gradoInstitucion int, ocupacion string, direccion string,
	tuvoTuberculosis bool,
	tieneInfTransSex bool,
	tieneDiabetes bool,
	tieneHta bool,
	tieneSobrepeso bool,
	tieneInfarto bool,
	tieneDislipenia bool,
	tieneInfRenalGlaucoma bool,
	tieneDepresionEsquizofrenia bool,
	antecedentes []string,
	tieneHospitaliacionTransfusiones bool,
	dispacidad []string,
	tieneConsumoTabaco bool,
	tieneConsumoAlcohol bool,
	tieneConsumoDrogas bool,
	tieneInterQuirurjica bool,
	cancer []string,
	tieneRiesgo bool,
	tieneViolencia bool,
	tieneSida bool,
	tieneITS bool,
	tieneHepatitis bool,
	tieneDbm bool,
	tieneCancer bool,
	tieneDepresion bool,
	tieneProbPsiquiatricos bool,
	otros []string,
	reaccionMedicamentos []string,
	medicamenteFrecuente []string,
	edadInicioRelacionSexual int,
	numParejas int,
	hijosVivos int,
	rsMismoSexo bool,
	menarquia interface{},
	flujoVagPatologico bool,
	dismenorrea bool,
	tieneEmbarazo bool,
	tieneParto bool,
	tienePrematuro bool,
	tieneAborto bool,
	gestacion interface{},
	tieneFiebre15Dias bool,
	tieneTos15Dias bool,
	lesionesGenitales []string,
	presionArterial interface{},
	tieneVacAntitetanica bool,
	tieneVacAntiamerilica bool,
	tieneVacAntihepatitisB bool,
	tieneEncias bool,
	tieneCaries bool,
	tieneEdentulismoParcial bool,
	tieneEdentulismoTotal bool,
	tieneUrgTratamientoBucal bool,
	tieneAnsiedad bool,
	tieneExamVisual bool,
	tieneExamColesterol bool,
	tieneExamGlucosa bool,
	tieneExamMamas bool,
	tieneExamProstata bool,
	tieneExamPelvicoPap bool,
	tieneExamMamografia bool,
	tieneHabFisica bool,
	tieneHabAlcohol bool, tieneHabDrogas bool, tienePlanificacionSexual bool) *HistoriaClinica {
	return &HistoriaClinica{
		Id:                               id,
		IdMedico:                         idMedico,
		IdPaciente:                       idPaciente,
		EstadoCivil:                      estadoCivil,
		GradoInstitucion:                 gradoInstitucion,
		Ocupacion:                        ocupacion,
		Direccion:                        direccion,
		IdDistrito:                       idDistrito,
		FechaRegistro:                    time.Now(),
		TuvoTuberculosis:                 tuvoTuberculosis,
		TieneInfTransSex:                 tieneInfTransSex,
		TieneDiabetes:                    tieneDiabetes,
		TieneHta:                         tieneHta,
		TieneSobrepeso:                   tieneSobrepeso,
		TieneInfarto:                     tieneInfarto,
		TieneDislipenia:                  tieneDislipenia,
		TieneInfRenalGlaucoma:            tieneInfRenalGlaucoma,
		TieneDepresionEsquizofrenia:      tieneDepresionEsquizofrenia,
		Antecedentes:                     antecedentes,
		TieneHospitaliacionTransfusiones: tieneHospitaliacionTransfusiones,
		Dispacidad:                       dispacidad,
		TieneConsumoTabaco:               tieneConsumoTabaco,
		TieneConsumoAlcohol:              tieneConsumoAlcohol,
		TieneConsumoDrogas:               tieneConsumoDrogas,
		TieneInterQuirurjica:             tieneInterQuirurjica,
		Cancer:                           cancer,
		TieneRiesgo:                      tieneRiesgo,
		TieneViolencia:                   tieneViolencia,
		TieneSida:                        tieneSida,
		TieneITS:                         tieneITS,
		TieneHepatitis:                   tieneHepatitis,
		TieneDbm:                         tieneDbm,
		TieneCancer:                      tieneCancer,
		TieneDepresion:                   tieneDepresion,
		TieneProbPsiquiatricos:           tieneProbPsiquiatricos,
		Otros:                            otros,
		ReaccionMedicamentos:             reaccionMedicamentos,
		MedicamenteFrecuente:             medicamenteFrecuente,
		EdadInicioRelacionSexual:         edadInicioRelacionSexual,
		NumParejas:                       numParejas,
		HijosVivos:                       hijosVivos,
		RsMismoSexo:                      rsMismoSexo,
		Menarquia:                        menarquia,
		FlujoVagPatologico:               flujoVagPatologico,
		Dismenorrea:                      dismenorrea,
		TieneEmbarazo:                    tieneEmbarazo,
		TieneParto:                       tieneParto,
		TienePrematuro:                   tienePrematuro,
		TieneAborto:                      tieneAborto,
		Gestacion:                        gestacion,
		TieneFiebre15Dias:                tieneFiebre15Dias,
		TieneTos15Dias:                   tieneTos15Dias,
		LesionesGenitales:                lesionesGenitales,
		PresionArterial:                  presionArterial,
		TieneVacAntitetanica:             tieneVacAntitetanica,
		TieneVacAntiamerilica:            tieneVacAntiamerilica,
		TieneVacAntihepatitisB:           tieneVacAntihepatitisB,
		TieneEncias:                      tieneEncias,
		TieneCaries:                      tieneCaries,
		TieneEdentulismoParcial:          tieneEdentulismoParcial,
		TieneEdentulismoTotal:            tieneEdentulismoTotal,
		TieneUrgTratamientoBucal:         tieneUrgTratamientoBucal,
		TieneAnsiedad:                    tieneAnsiedad,
		TieneExamVisual:                  tieneExamVisual,
		TieneExamColesterol:              tieneExamColesterol,
		TieneExamGlucosa:                 tieneExamGlucosa,
		TieneExamMamas:                   tieneExamMamas,
		TieneExamProstata:                tieneExamProstata,
		TieneExamPelvicoPap:              tieneExamPelvicoPap,
		TieneExamMamografia:              tieneExamMamografia,
		TieneHabFisica:                   tieneHabFisica,
		TieneHabAlcohol:                  tieneHabAlcohol,
		TieneHabDrogas:                   tieneHabDrogas,
		TienePlanificacionSexual:         tienePlanificacionSexual,
	}
}

func UpdateHistoriaClinica(
	id string,
	idPaciente string, estadoCivil int, gradoInstitucion int, ocupacion string, direccion string,
	idDistrito int,
	tuvoTuberculosis bool,
	tieneInfTransSex bool,
	tieneDiabetes bool,
	tieneHta bool,
	tieneSobrepeso bool,
	tieneInfarto bool,
	tieneDislipenia bool,
	tieneInfRenalGlaucoma bool,
	tieneDepresionEsquizofrenia bool,
	antecedentes []string,
	tieneHospitaliacionTransfusiones bool,
	dispacidad []string,
	tieneConsumoTabaco bool,
	tieneConsumoAlcohol bool,
	tieneConsumoDrogas bool,
	tieneInterQuirurjica bool,
	cancer []string,
	tieneRiesgo bool,
	tieneViolencia bool,
	tieneSida bool,
	tieneITS bool,
	tieneHepatitis bool,
	tieneDbm bool,
	tieneCancer bool,
	tieneDepresion bool,
	tieneProbPsiquiatricos bool,
	otros []string,
	reaccionMedicamentos []string,
	medicamenteFrecuente []string,
	edadInicioRelacionSexual int,
	numParejas int,
	hijosVivos int,
	rsMismoSexo bool,
	menarquia interface{},
	flujoVagPatologico bool,
	dismenorrea bool,
	tieneEmbarazo bool,
	tieneParto bool,
	tienePrematuro bool,
	tieneAborto bool,
	gestacion interface{},
	tieneFiebre15Dias bool,
	tieneTos15Dias bool,
	lesionesGenitales []string,
	presionArterial interface{},
	tieneVacAntitetanica bool,
	tieneVacAntiamerilica bool,
	tieneVacAntihepatitisB bool,
	tieneEncias bool,
	tieneCaries bool,
	tieneEdentulismoParcial bool,
	tieneEdentulismoTotal bool,
	tieneUrgTratamientoBucal bool,
	tieneAnsiedad bool,
	tieneExamVisual bool,
	tieneExamColesterol bool,
	tieneExamGlucosa bool,
	tieneExamMamas bool,
	tieneExamProstata bool,
	tieneExamPelvicoPap bool,
	tieneExamMamografia bool,
	tieneHabFisica bool,
	tieneHabAlcohol bool,
	tieneHabDrogas bool,
	tienePlanificacionSexual bool) *HistoriaClinica {
	return &HistoriaClinica{
		Id:                               id,
		IdPaciente:                       idPaciente,
		EstadoCivil:                      estadoCivil,
		GradoInstitucion:                 gradoInstitucion,
		Ocupacion:                        ocupacion,
		Direccion:                        direccion,
		IdDistrito:                       idDistrito,
		TuvoTuberculosis:                 tuvoTuberculosis,
		TieneInfTransSex:                 tieneInfTransSex,
		TieneDiabetes:                    tieneDiabetes,
		TieneHta:                         tieneHta,
		TieneSobrepeso:                   tieneSobrepeso,
		TieneInfarto:                     tieneInfarto,
		TieneDislipenia:                  tieneDislipenia,
		TieneInfRenalGlaucoma:            tieneInfRenalGlaucoma,
		TieneDepresionEsquizofrenia:      tieneDepresionEsquizofrenia,
		Antecedentes:                     antecedentes,
		TieneHospitaliacionTransfusiones: tieneHospitaliacionTransfusiones,
		Dispacidad:                       dispacidad,
		TieneConsumoTabaco:               tieneConsumoTabaco,
		TieneConsumoAlcohol:              tieneConsumoAlcohol,
		TieneConsumoDrogas:               tieneConsumoDrogas,
		TieneInterQuirurjica:             tieneInterQuirurjica,
		Cancer:                           cancer,
		TieneRiesgo:                      tieneRiesgo,
		TieneViolencia:                   tieneViolencia,
		TieneSida:                        tieneSida,
		TieneITS:                         tieneITS,
		TieneHepatitis:                   tieneHepatitis,
		TieneDbm:                         tieneDbm,
		TieneCancer:                      tieneCancer,
		TieneDepresion:                   tieneDepresion,
		TieneProbPsiquiatricos:           tieneProbPsiquiatricos,
		Otros:                            otros,
		ReaccionMedicamentos:             reaccionMedicamentos,
		MedicamenteFrecuente:             medicamenteFrecuente,
		EdadInicioRelacionSexual:         edadInicioRelacionSexual,
		NumParejas:                       numParejas,
		HijosVivos:                       hijosVivos,
		RsMismoSexo:                      rsMismoSexo,
		Menarquia:                        menarquia,
		FlujoVagPatologico:               flujoVagPatologico,
		Dismenorrea:                      dismenorrea,
		TieneEmbarazo:                    tieneEmbarazo,
		TieneParto:                       tieneParto,
		TienePrematuro:                   tienePrematuro,
		TieneAborto:                      tieneAborto,
		Gestacion:                        gestacion,
		TieneFiebre15Dias:                tieneFiebre15Dias,
		TieneTos15Dias:                   tieneTos15Dias,
		LesionesGenitales:                lesionesGenitales,
		PresionArterial:                  presionArterial,
		TieneVacAntitetanica:             tieneVacAntitetanica,
		TieneVacAntiamerilica:            tieneVacAntiamerilica,
		TieneVacAntihepatitisB:           tieneVacAntihepatitisB,
		TieneEncias:                      tieneEncias,
		TieneCaries:                      tieneCaries,
		TieneEdentulismoParcial:          tieneEdentulismoParcial,
		TieneEdentulismoTotal:            tieneEdentulismoTotal,
		TieneUrgTratamientoBucal:         tieneUrgTratamientoBucal,
		TieneAnsiedad:                    tieneAnsiedad,
		TieneExamVisual:                  tieneExamVisual,
		TieneExamColesterol:              tieneExamColesterol,
		TieneExamGlucosa:                 tieneExamGlucosa,
		TieneExamMamas:                   tieneExamMamas,
		TieneExamProstata:                tieneExamProstata,
		TieneExamPelvicoPap:              tieneExamPelvicoPap,
		TieneExamMamografia:              tieneExamMamografia,
		TieneHabFisica:                   tieneHabFisica,
		TieneHabAlcohol:                  tieneHabAlcohol,
		TieneHabDrogas:                   tieneHabDrogas,
		TienePlanificacionSexual:         tienePlanificacionSexual,
	}
}
