package routes

import (
	"log"
	"net/http"

	"goAPI/controllers"
	"goAPI/metric"
	"goAPI/middleware"
	"goAPI/shared"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	shared.ZapLogCustom([]string{"Criando rotas."}, "info")

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades/todas", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/unidade/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/add", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/del/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/upgrade/{id}", controllers.EditaPersonalidade).Methods("Put")

	shared.ZapLogCustom([]string{"Iniciando a API."}, "info")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

func RunMetrics() {

	metricService, _ := metric.NewPrometheusService()
	g := gin.New()
	gin.SetMode("release")
	g.Use(middleware.Metrics(metricService))
	g.GET("/metrics", func(c *gin.Context) {
		handler := promhttp.Handler()
		handler.ServeHTTP(c.Writer, c.Request)
	})
	g.Run(":9990")

}
