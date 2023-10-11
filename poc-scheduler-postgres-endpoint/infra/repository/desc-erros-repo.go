package repository

import (
	domain "poc-scheduler-postgres-endpoint/domain/database"
)

type IDescricaoErrosRepository interface {
	Recuperar(idCte string) ([]domain.DescricaoErros, error)
}
