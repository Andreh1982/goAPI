package main

import (
	"goAPI/database"
	"goAPI/routes"
	"goAPI/shared"
)

func main() {

	go routes.RunMetrics()

	shared.ZapLogCustom([]string{"Iniciando o servidor goAPI"}, "info")

	database.ConectaComBancoDeDados()

	routes.HandleResquest()

}
