package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	middleware_api "github.com/UPC-Works/api-aliviate/api/middlewares"
	admin_service "github.com/UPC-Works/api-aliviate/services/admin"
	analisis_historia_service "github.com/UPC-Works/api-aliviate/services/analisis_historia"
	analisis_laboratorio_campo_service "github.com/UPC-Works/api-aliviate/services/analisis_laboratorio_campo"
	analisis_laboratorio_codigo_service "github.com/UPC-Works/api-aliviate/services/analisis_laboratorio_codigo"
	auth_service "github.com/UPC-Works/api-aliviate/services/auth"
	consulta_service "github.com/UPC-Works/api-aliviate/services/consulta"
	documento_historia_service "github.com/UPC-Works/api-aliviate/services/documento_historia"
	establecimiento_service "github.com/UPC-Works/api-aliviate/services/establecimiento"
	historia_clinica_service "github.com/UPC-Works/api-aliviate/services/historia_clinica"
	medico_service "github.com/UPC-Works/api-aliviate/services/medico"
	paciente_service "github.com/UPC-Works/api-aliviate/services/paciente"
	prediccion_service "github.com/UPC-Works/api-aliviate/services/prediccion"
)

func HandlerRouters() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", index)

	//V1
	version_1 := e.Group("/v1")

	//V1 - AUTH
	router_auth := version_1.Group("/auth")
	router_auth.GET("", auth_service.Authentication)

	//V1 - ADMIN
	router_admin := version_1.Group("/admin")
	router_admin.POST("/sign-up", admin_service.SignUp)
	router_admin.POST("/login", admin_service.Login)
	router_admin.POST("/assign-est-medico", admin_service.AsignarEstablecimiento)
	router_admin.POST("/create-analisis", analisis_laboratorio_codigo_service.Add)
	router_admin.GET("/list-analisis", analisis_laboratorio_codigo_service.GetAll)
	router_admin.POST("/create-analisis-campo", analisis_laboratorio_campo_service.Add)
	router_admin.GET("/list-analisis-campo", analisis_laboratorio_campo_service.GetAll)

	//V1 - ESTABLECIMIENTO
	router_establecimiento := version_1.Group("/establecimiento", middleware_api.Auth)
	router_establecimiento.POST("", establecimiento_service.Add)
	router_establecimiento.GET("", establecimiento_service.GetAll)

	//V1 - CONSULTA
	router_consulta := version_1.Group("/consulta", middleware_api.Auth)
	router_consulta.POST("", consulta_service.Add)
	router_consulta.GET("", consulta_service.GetAll)

	//V1 - HISTORIA CLINICA
	router_historia := version_1.Group("/historia_clinica", middleware_api.Auth)
	router_historia.POST("", historia_clinica_service.Add)
	router_historia.GET("", historia_clinica_service.GetAll)
	router_historia.GET("/modificaciones", historia_clinica_service.GetAll)
	router_historia.PUT("", historia_clinica_service.Update)
	router_historia.GET("/:id_historia_clinica", historia_clinica_service.GetOne)

	//V1 - MEDICO
	router_medico := version_1.Group("/medico")
	router_medico.POST("/sign-up", medico_service.SignUp)
	router_medico.POST("/login", medico_service.Login)
	router_medico.GET("", medico_service.GetAll, middleware_api.Auth)
	router_medico.PUT("", medico_service.Update, middleware_api.Auth)

	//V1 - PACIENTE
	router_paciente := version_1.Group("/paciente", middleware_api.Auth)
	router_paciente.POST("", paciente_service.Add)
	router_paciente.GET("", paciente_service.GetAll)
	router_paciente.PUT("", paciente_service.Update)

	//V1 - ANALISIS LABORATORIO
	router_analisis_laboratorio := version_1.Group("/analisis_laboratorio", middleware_api.Auth)
	router_analisis_laboratorio.GET("/list-analisis", analisis_laboratorio_codigo_service.GetAll)
	router_analisis_laboratorio.GET("/list-analisis-campo", analisis_laboratorio_campo_service.GetAll)
	router_analisis_laboratorio.POST("/register-analisis-historia", analisis_historia_service.Add)
	router_analisis_laboratorio.POST("/list-analisis-historia", analisis_historia_service.GetAll)

	//V1 - DOCUMENTOS
	router_documentos := version_1.Group("/documentos", middleware_api.Auth)
	router_documentos.POST("/:idhistoriaclinica", documento_historia_service.Add)
	router_documentos.GET("", documento_historia_service.GetAll)

	//V1 - PREDICCION
	router_prediccion := version_1.Group("/prediccion")
	router_prediccion.POST("/prediccion", prediccion_service.Predecir)

	//Open the port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "6500"
	}
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func index(c echo.Context) error {
	return c.JSON(401, "Unauthorized access")
}
