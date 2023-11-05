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
	analisis_laboratorio_service "github.com/UPC-Works/api-aliviate/services/analisis_laboratorio"
	auth_service "github.com/UPC-Works/api-aliviate/services/auth"
	consulta_service "github.com/UPC-Works/api-aliviate/services/consulta"
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
	router_analisis_laboratorio.POST("", analisis_laboratorio_service.Add)
	router_analisis_laboratorio.PUT("", analisis_laboratorio_service.Update)
	router_analisis_laboratorio.GET("", analisis_laboratorio_service.GetAll)

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
