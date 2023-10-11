package usecase

import (
	"github.com/segmentio/kafka-go"
)

type IMessagerConsumerUseCase interface {
	Executa(mensagem kafka.Message)
}
