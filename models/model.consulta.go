package models

import "time"

// Model

type Consulta struct {
	Id                            string    `json:"id"`
	IdHistoriaClinica             string    `json:"idHistoriaClinica"`
	IdMedico                      string    `json:"idMedico"`
	NombreMedico                  string    `json:"nombreMedico"`
	FechaRegistro                 time.Time `json:"fechaRegistro"`
	DescripcionEnfermedadPaciente string    `json:"descripcionEnfermedadPaciente"`
	TiempoEnfermedad              int       `json:"tiempoEnfermedad"`
	Apetito                       string    `json:"apetito"`
	Sed                           string    `json:"sed"`
	Suenio                        string    `json:"suenio"`
	EstadoAnimo                   string    `json:"estadoAnimo"`
	OtroDetalle                   string    `json:"otroDetalle"`
	Orina                         string    `json:"orina"`
	Deposiciones                  string    `json:"deposiciones"`
	Temperatura                   float32   `json:"temperatura"`
	PA                            float32   `json:"pA"`
	FC                            float32   `json:"FC"`
	FR                            float32   `json:"FR"`
	Peso                          float32   `json:"peso"`
	Talla                         float32   `json:"talla"`
	IMC                           float32   `json:"iMC"`
	Diagnostico                   string    `json:"diagnostico"`
	Tratamiento                   string    `json:"tratamiento"`
	DiagnosticoIA                 string    `json:"diagnosticoIA"`
	TratamientoIA                 string    `json:"tratamientoIA"`
	ExamenesAuxiliares            []string  `json:"examenesAuxiliares"`
	ProximaCita                   time.Time `json:"proximaCita"`
	Observaciones                 string    `json:"observaciones"`
	NombreCompletoPaciente        string    `json:"nombreCompletoPaciente"`
	IdPaciente                    string    `json:"idPaciente"`
	SignosSintomas                string    `json:"signosSintomas"`
}

//Constructor

func NewConsulta(id string, idHistoriaClinica string, idMedico string, descripcionEnfermedadPaciente string, tiempoEnfermedad int, apetito string, sed string, suenio string, estadoAnimo string, otroDetalle string, orina string, deposiciones string, temperatura float32, pA float32, fC float32, fR float32, peso float32, talla float32, iMC float32, diagnostico string, tratamiento string, diagnosticoIA string, tratamientoIA string, examenesAuxiliares []string, proximaCita time.Time, observaciones string, signosSintomas string) *Consulta {
	return &Consulta{
		Id:                            id,
		IdHistoriaClinica:             idHistoriaClinica,
		IdMedico:                      idMedico,
		FechaRegistro:                 time.Now(),
		DescripcionEnfermedadPaciente: descripcionEnfermedadPaciente,
		TiempoEnfermedad:              tiempoEnfermedad,
		Apetito:                       apetito,
		Sed:                           sed,
		Suenio:                        suenio,
		EstadoAnimo:                   estadoAnimo,
		OtroDetalle:                   otroDetalle,
		Orina:                         orina,
		Deposiciones:                  deposiciones,
		Temperatura:                   temperatura,
		PA:                            pA,
		FC:                            fC,
		FR:                            fR,
		Peso:                          peso,
		Talla:                         talla,
		IMC:                           iMC,
		Diagnostico:                   diagnostico,
		Tratamiento:                   tratamiento,
		DiagnosticoIA:                 diagnosticoIA,
		TratamientoIA:                 tratamientoIA,
		ExamenesAuxiliares:            examenesAuxiliares,
		ProximaCita:                   proximaCita,
		Observaciones:                 observaciones,
		SignosSintomas:                signosSintomas,
	}
}

func UpdateConsulta(id string, idHistoriaClinica string, descripcionEnfermedadPaciente string, tiempoEnfermedad int, apetito string, sed string, suenio string, estadoAnimo string, otroDetalle string, orina string, deposiciones string, temperatura float32, pA float32, fC float32, fR float32, peso float32, talla float32, iMC float32, diagnostico string, tratamiento string, diagnosticoIA string, tratamientoIA string, examenesAuxiliares []string, proximaCita time.Time, observaciones string, signosSintomas string) *Consulta {
	return &Consulta{
		Id:                            id,
		IdHistoriaClinica:             idHistoriaClinica,
		DescripcionEnfermedadPaciente: descripcionEnfermedadPaciente,
		TiempoEnfermedad:              tiempoEnfermedad,
		Apetito:                       apetito,
		Sed:                           sed,
		Suenio:                        suenio,
		EstadoAnimo:                   estadoAnimo,
		OtroDetalle:                   otroDetalle,
		Orina:                         orina,
		Deposiciones:                  deposiciones,
		Temperatura:                   temperatura,
		PA:                            pA,
		FC:                            fC,
		FR:                            fR,
		Peso:                          peso,
		Talla:                         talla,
		IMC:                           iMC,
		Diagnostico:                   diagnostico,
		Tratamiento:                   tratamiento,
		DiagnosticoIA:                 diagnosticoIA,
		TratamientoIA:                 tratamientoIA,
		ExamenesAuxiliares:            examenesAuxiliares,
		ProximaCita:                   proximaCita,
		Observaciones:                 observaciones,
		SignosSintomas:                signosSintomas,
	}
}
