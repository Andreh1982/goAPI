package routes

import (
	"goAPI/controllers"
	"goAPI/metric"
	"goAPI/middleware"

	"goAPI/shared"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleRequest() {

	router := gin.Default()
	gin.SetMode("release")

	shared.ZapLogCustom([]string{"Inicializando métricas."}, "info")

	metricService, _ := metric.NewPrometheusService()
	router.Use(middleware.Metrics(metricService))
	router.GET("/metrics", func(c *gin.Context) {
		promHandler := promhttp.Handler()
		promHandler.ServeHTTP(c.Writer, c.Request)
	})

	shared.ZapLogCustom([]string{"Criando rotas."}, "info")

	router.GET("/", controllers.Home)
	router.GET("/api/personalidades/todas", controllers.TodasPersonalidades)
	router.GET("/api/personalidades/unidade/:id", controllers.RetornaUmaPersonalidade)
	router.POST("/api/personalidades/add", controllers.CriaUmaNovaPersonalidade)
	router.PUT("/api/personalidades/update/:id", controllers.EditaPersonalidade)
	router.DELETE("/api/personalidades/del/:id", controllers.DeletaUmaPersonalidade)

	shared.ZapLogCustom([]string{"Iniciando a API."}, "info")

	router.Run(":9990")

}
