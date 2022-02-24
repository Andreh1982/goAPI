package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"goAPI/database"
	"goAPI/models"
	"goAPI/shared"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
	shared.ZapLogCustom([]string{"Retornando todas as entradas"}, "info")

	size := "Total de entradas: " + strconv.Itoa(len(p))
	shared.ZapLogCustom([]string{size}, "info")

}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

	shared.ZapLogCustom([]string{personalidade.Nome}, "info")
	shared.ZapLogCustom([]string{personalidade.Historia}, "info")

}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)

	shared.ZapLogCustom([]string{"Criando nova entrada..."}, "info")
	shared.ZapLogCustom([]string{novaPersonalidade.Nome}, "info")
	shared.ZapLogCustom([]string{novaPersonalidade.Historia}, "info")
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	shared.ZapLogCustom([]string{"[DEL] Deletando entrada..."}, "info")
	shared.ZapLogCustom([]string{personalidade.Nome}, "info")
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
