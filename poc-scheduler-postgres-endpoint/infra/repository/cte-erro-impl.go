package repository

import (
	"fmt"
	"time"

	domain "poc-scheduler-postgres-endpoint/domain/database"
	"poc-scheduler-postgres-endpoint/domain/log"
	"poc-scheduler-postgres-endpoint/infra/database"
)

const (
	filtrochaveCteError    = "status=?"
	filtroChaveAtualizacao = "id_cte=?"
)

var (
	ctesError     []domain.CteErro = []domain.CteErro{}
	status        string           = "E"
	statusVencido string           = "V"
	dataAlteracao time.Time        = time.Now()
)

type CteErroRepoPg struct {
}

func NewCteErroPgRepository() ICteErroRepository {
	return &CteErroRepoPg{}
}

func (r *CteErroRepoPg) Recuperar() ([]domain.CteErro, error) {
	err := database.DB.Where(
		filtrochaveCteError,
		status).Find(&ctesError).Error
	if err != nil && err.Error() != "record not found" {
		log.GravarLog(fmt.Sprintf("Deu erro no find: %s", err.Error()))
		return ctesError, err
	}

	if err != nil && err.Error() == "record not found" {
		return ctesError, nil
	}

	return ctesError, nil
}

func (r *CteErroRepoPg) AtualizaStatus(ctesErrorVencidas []domain.CteErro) error {
	for _, cte := range ctesErrorVencidas {
		err := database.DB.Model(&domain.CteErro{}).Where(filtroChaveAtualizacao, cte.IdCte).Updates(map[string]interface{}{
			"status":       statusVencido,
			"dt_alteracao": dataAlteracao,
		}).Error
		if err != nil {
			log.GravarLog(err.Error())
			return err
		}
	}
	return nil
}
