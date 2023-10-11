package consumer

import (
	"poc-consumer-go-endpoint-integrador/usecase"

	"github.com/segmentio/kafka-go"
)

type IKafkaConsumer interface {
	ConsumirMensagem()
}

type KafkaConsumerImpl struct {
}

var (
	canal        chan kafka.Message
	businesscase usecase.IMessagerConsumerUseCase
)

func NewKafkaConsumer(canalmsg chan kafka.Message, businesscaseinput usecase.IMessagerConsumerUseCase) IKafkaConsumer {
	canal = canalmsg
	businesscase = businesscaseinput
	return &KafkaConsumerImpl{}
}

func (*KafkaConsumerImpl) ConsumirMensagem() {
	for msg := range canal {
		businesscase.Executa(msg)
	}
}
