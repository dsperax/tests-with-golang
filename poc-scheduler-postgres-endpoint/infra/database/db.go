package database

import (
	"fmt"
	"os"

	"poc-scheduler-postgres-endpoint/app/utils"
	"poc-scheduler-postgres-endpoint/domain/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := utils.GetVariavelAmbiente("pg_auditoria_fretes_secret", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s", "localhost", "postgres", "root", "postgres", 5432, "disable"))
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.GravarLog("Erro ao conectar com o banco de dados! Descrição do erro:" + err.Error())
		os.Exit(1)
	}
}
