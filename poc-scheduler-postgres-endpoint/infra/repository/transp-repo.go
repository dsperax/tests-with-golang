package repository

import (
	domain "poc-scheduler-postgres-endpoint/domain/database"
)

type ITransportadorRepository interface {
	Recuperar(cnpj string) (domain.Transportador, error)
}
