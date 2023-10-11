package repository

import (
	"fmt"

	domain "poc-scheduler-postgres-endpoint/domain/database"
	"poc-scheduler-postgres-endpoint/domain/log"
	"poc-scheduler-postgres-endpoint/infra/database"
)

const (
	filtrochaveDescricaoErros = "id_cte=?"
)

var (
	descErros []domain.DescricaoErros = []domain.DescricaoErros{}
)

type DescricaoErrosRepoPg struct {
}

func NewDescricaoErrosPgRepository() IDescricaoErrosRepository {
	return &DescricaoErrosRepoPg{}
}

func (r *DescricaoErrosRepoPg) Recuperar(idCte string) ([]domain.DescricaoErros, error) {
	err := database.DB.Where(
		filtrochaveDescricaoErros,
		idCte).Find(&descErros).Error
	if err != nil && err.Error() != "record not found" {
		log.GravarLog(fmt.Sprintf("Deu erro no find: %s", err.Error()))
		return descErros, err
	}

	if err != nil && err.Error() == "record not found" {
		return descErros, nil
	}

	return descErros, nil
}
