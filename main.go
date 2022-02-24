package main

import (
	"goAPI/database"
	"goAPI/routes"
	"goAPI/shared"
)

func main() {

	shared.ZapLogCustom("Iniciando o servidor goAPI", "info")

	database.ConectaComBancoDeDados()

	routes.HandleResquest()

}
