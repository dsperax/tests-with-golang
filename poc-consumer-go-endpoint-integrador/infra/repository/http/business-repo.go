package repository

import (
	domain "poc-consumer-go-endpoint-integrador/domain/message"
)

type IBusinessRepository interface {
	EndPointRegras(businessdomain domain.BusinessMessage) error
}
