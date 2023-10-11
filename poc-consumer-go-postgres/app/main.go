package main

import (
	"strings"

	"poc-consumer-go-postgres/app/router"
	"poc-consumer-go-postgres/app/utils"
	"poc-consumer-go-postgres/infra/database"
	akafka "poc-consumer-go-postgres/infra/kafka/config"
	"poc-consumer-go-postgres/infra/kafka/consumer"
	"poc-consumer-go-postgres/infra/repository"
	"poc-consumer-go-postgres/usecase"

	"github.com/segmentio/kafka-go"
)

var (
	canal           chan kafka.Message             = make(chan kafka.Message)
	canalresultado  chan bool                      = make(chan bool)
	businessrepo    repository.IBusinessRepository = repository.NewBusinessMSSQLRepository()
	businessusecase usecase.IKafkaConsumerUseCase  = usecase.NewBusinessUseCase(businessrepo, canalresultado)
	consumerimpl    consumer.IKafkaConsumer        = consumer.NewKafkaConsumer(canal, businessusecase)
)

func main() {
	database.ConectaComBancoDeDados()

	topico := utils.GetVariavelAmbiente("kafka_topic", "topico-posicao")
	servidores := strings.Split(utils.GetVariavelAmbiente("kafka_servers", "localhost:9092"), ";")
	groupid := utils.GetVariavelAmbiente("kafka_groupid", "consumer-dev")

	go akafka.Consumer(topico, servidores, groupid, canal, canalresultado)
	go consumerimpl.ConsumirMensagem()

	router.Roteador()
}
