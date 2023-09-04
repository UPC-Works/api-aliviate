package helper

type ErrorStructure struct {
	HasError bool   `json:"hasError"`
	Detail   string `json:"detail"`
}

type JwtStructure struct {
	JWT            string `json:"jwt"`
	NombreCompleto string `json:"nombreCompleto"`
	Correo         string `json:"correo"`
	Rol            int    `json:"rol"`
}

type AuthStructure struct {
	Id             string `json:"id"`
	NombreCompleto string `json:"nombreCompleto"`
	Correo         string `json:"correo"`
	Rol            int    `json:"rol"`
}
