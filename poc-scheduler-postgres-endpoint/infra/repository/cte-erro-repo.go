package repository

import (
	domain "poc-scheduler-postgres-endpoint/domain/database"
)

type ICteErroRepository interface {
	Recuperar() ([]domain.CteErro, error)
	AtualizaStatus(ctesErrorVencidas []domain.CteErro) error
}
