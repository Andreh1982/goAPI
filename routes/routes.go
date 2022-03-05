package routes

import (
	"goAPI/controllers"
	"goAPI/database"
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

	api := &controllers.APIEnv{
		DB: database.GetDB(),
	}

	router.GET("", api.GetPersons)
	router.GET("/:id", api.GetPerson)
	router.POST("", api.CreatePerson)
	router.PUT("/:id", api.UpdatePerson)
	router.DELETE("/:id", api.DeletePerson)

	shared.ZapLogCustom([]string{"Iniciando a API."}, "info")

	router.Run(":9990")

}
