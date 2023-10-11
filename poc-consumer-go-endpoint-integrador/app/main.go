package main

import (
	"strings"

	"poc-consumer-go-endpoint-integrador/app/router"
	"poc-consumer-go-endpoint-integrador/app/utils"
	akafka "poc-consumer-go-endpoint-integrador/infra/message/config"
	"poc-consumer-go-endpoint-integrador/infra/message/consumer"
	repository "poc-consumer-go-endpoint-integrador/infra/repository/http"
	"poc-consumer-go-endpoint-integrador/usecase"

	"github.com/segmentio/kafka-go"
)

var (
	canal           chan kafka.Message               = make(chan kafka.Message)
	canalresultado  chan bool                        = make(chan bool)
	businessrepo    repository.IBusinessRepository   = repository.NewBusinessMSSQLRepository()
	businessusecase usecase.IMessagerConsumerUseCase = usecase.NewBusinessUseCase(businessrepo, canalresultado)
	consumerimpl    consumer.IKafkaConsumer          = consumer.NewKafkaConsumer(canal, businessusecase)
)

func main() {
	topico := utils.GetVariavelAmbiente("kafka_topic", "topico-posicao")
	servidores := strings.Split(utils.GetVariavelAmbiente("kafka_servers", "localhost:9092"), ";")
	groupid := utils.GetVariavelAmbiente("kafka_groupid", "consumer-dev")

	go akafka.Consumer(topico, servidores, groupid, canal, canalresultado)
	go consumerimpl.ConsumirMensagem()

	router.Roteador()
}
