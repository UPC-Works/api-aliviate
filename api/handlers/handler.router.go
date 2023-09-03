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
	establecimiento_service "github.com/UPC-Works/api-aliviate/services/establecimiento"
	historia_clinica_service "github.com/UPC-Works/api-aliviate/services/historia_clinica"
	medico_service "github.com/UPC-Works/api-aliviate/services/medico"
	paciente_service "github.com/UPC-Works/api-aliviate/services/paciente"
)

func HandlerRouters() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/", index)

	//V1
	version_1 := e.Group("/v1")

	//V1 - ADMIN
	router_admin := version_1.Group("/admin")
	router_admin.POST("/sign-up", admin_service.SignUp)
	router_admin.POST("/login", admin_service.Login)

	//V1 - ESTABLECIMIENTO
	router_establecimiento := version_1.Group("/establecimiento", middleware_api.Auth)
	router_establecimiento.POST("", establecimiento_service.Add)
	/*router_warehouse.GET("", warehouse_service.GetAll)
	router_warehouse.PUT("", warehouse_service.Update)
	router_warehouse.DELETE("", warehouse_service.SendToDelete)*/

	//V1 - HISTORIA CLINICA
	router_historia := version_1.Group("/historia_clinica", middleware_api.Auth)
	router_historia.POST("", historia_clinica_service.Add)
	/*router_provider.GET("", provider_service.GetAll)
	router_provider.PUT("", provider_service.Update)
	router_provider.DELETE("", provider_service.SendToDelete)*/

	//V1 - MEDICO
	router_medico := version_1.Group("/medico")
	router_medico.POST("/sign-up", medico_service.SignUp)
	router_medico.POST("/login", medico_service.Login)
	/*router_medico.GET("", measure_service.GetAll)
	router_medico.PUT("", provider_service.Update)
	router_medico.DELETE("", measure_service.Delete)*/

	//V1 - PACIENTE
	router_paciente := version_1.Group("/paciente", middleware_api.Auth)
	router_paciente.POST("", paciente_service.Add)
	/*router_supply.GET("", supply_service.GetAll)
	router_supply.PUT("", supply_service.Update)
	router_supply.DELETE("", supply_service.SendToDelete)*/

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
