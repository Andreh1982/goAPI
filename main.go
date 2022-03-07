package main

import (
	"goAPI/database"
	"goAPI/routes"
	"goAPI/shared"
)

func main() {

	shared.ZapLogCustom([]string{"Iniciando o servidor goAPI"}, "info")

	database.ConnectDB()

	routes.HandleRequest()

}
