package repository

import (
	domain "poc-consumer-go-mongo-async/domain/database"
)

type IBusinessRepository interface {
	ExisteRegistro(businessdomain domain.BusinessEntity) (bool, error)
	Inserir(businessdomain domain.BusinessEntity) error
	Atualizar(businessdomain domain.BusinessEntity) error
	Excluir(businessdomain domain.BusinessEntity) error
}
