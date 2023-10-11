package repository

import (
	"fmt"

	domain "poc-scheduler-postgres-endpoint/domain/database"
	"poc-scheduler-postgres-endpoint/domain/log"
	"poc-scheduler-postgres-endpoint/infra/database"
)

const (
	filtrochaveTransportador = "cd_cnpj_transportador=?"
)

var (
	transportador domain.Transportador
)

type TransportadorRepoPg struct {
}

func NewTransportadorPgRepository() ITransportadorRepository {
	return &TransportadorRepoPg{}
}

func (r *TransportadorRepoPg) Recuperar(cnpj string) (domain.Transportador, error) {
	err := database.DB.Where(
		filtrochaveTransportador,
		cnpj).First(&transportador).Error
	if err != nil && err.Error() != "record not found" {
		log.GravarLog(fmt.Sprintf("Deu erro no find: %s", err.Error()))
		return transportador, err
	}

	if err != nil && err.Error() == "record not found" {
		return transportador, nil
	}
	return transportador, nil
}
