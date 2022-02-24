package routes

import (
	"log"
	"net/http"

	"goAPI/controllers"
	"goAPI/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"goAPI/shared"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	shared.ZapLogCustom("Criando rotas.", "info")

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades/todas", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/unidade/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/add", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/del/{id}", controllers.DeletaUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/upgrade/{id}", controllers.EditaPersonalidade).Methods("Put")

	shared.ZapLogCustom("Iniciando API.", "info")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
