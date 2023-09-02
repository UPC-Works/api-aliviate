package helper

type ErrorStructure struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type JwtStructure struct {
	JWT      string `json:"jwt"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
}

type AuthStructure struct {
	Id       string `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Correo   string `json:"correo"`
}
