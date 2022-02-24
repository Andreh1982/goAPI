package database

import (
	"goAPI/shared"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		shared.ZapLogCustom("Erro ao conectar ao banco de dados!", "fatal")
		os.Exit(1)
	} else {
		shared.ZapLogCustom("Conectado ao banco de dados.", "info")
	}
}
