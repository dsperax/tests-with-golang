package main

import (
	"strconv"
	"strings"

	"poc-consumer-go-mongo-sync/app/router"
	"poc-consumer-go-mongo-sync/app/utils"
	"poc-consumer-go-mongo-sync/infra/database"
	akafka "poc-consumer-go-mongo-sync/infra/kafka/config"
	"poc-consumer-go-mongo-sync/infra/kafka/consumer"
	"poc-consumer-go-mongo-sync/infra/repository"
	"poc-consumer-go-mongo-sync/usecase"

	"github.com/segmentio/kafka-go"
)

var (
	canal           chan kafka.Message             = make(chan kafka.Message)
	canalresultado  chan bool                      = make(chan bool)
	businessrepo    repository.IBusinessRepository = repository.NewBusinessMongoDbRepository()
	businessusecase usecase.IKafkaConsumerUseCase  = usecase.NewBusinessUseCase(businessrepo, canalresultado)
	consumerimpl    consumer.IKafkaConsumer        = consumer.NewKafkaConsumer(canal, businessusecase)
)

func main() {
	database.ConectaComBancoDeDados()

	topico := utils.GetVariavelAmbiente("kafka_topic", "topico-posicao")
	servidores := strings.Split(utils.GetVariavelAmbiente("kafka_servers", "localhost:9092"), ";")
	groupid := utils.GetVariavelAmbiente("kafka_groupid", "consumer-dev-mongo")
	threadsstr := utils.GetVariavelAmbiente("kafka_threads", "1")
	threads, err := strconv.Atoi(threadsstr)
	if err != nil {
		threads = 1
	}

	for i := 0; i < threads; i++ {
		go akafka.Consumer(i, topico, servidores, groupid, canal, canalresultado)
	}

	go consumerimpl.ConsumirMensagem()

	router.Roteador()
}
