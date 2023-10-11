package akafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"poc-consumer-go-mongo-sync/app/utils"
	"poc-consumer-go-mongo-sync/domain/log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

var worker int

func Consumer(id int, topic string, servers []string, groupid string, mensagem chan kafka.Message, resultado chan bool) {
	for {
		CriaConsumer(id, topic, servers, groupid, mensagem, resultado)
	}
}

func CriaConsumer(id int, topic string, servers []string, groupid string, mensagem chan kafka.Message, resultado chan bool) {
	kafkaConsumer := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           servers,
		Topic:             topic,
		GroupID:           groupid,
		HeartbeatInterval: time.Second * 30,
		Dialer:            kafkaAuthentication(),
		// Logger:            kafka.LoggerFunc(logf),
		// ErrorLogger:       kafka.LoggerFunc(logf),
	})

	ctx := context.Background()

	for {
		msg, err := kafkaConsumer.FetchMessage(ctx)
		if err != nil {
			log.GravarLog(fmt.Sprintf("FetchMessage Erro: %s", err.Error()))
			kafkaConsumer.Close()
			return
		}

		mensagem <- msg

		if <-resultado {
			kafkaConsumer.CommitMessages(ctx, msg)
		} else {
			kafkaConsumer.Close()
			return
		}
	}
}

func kafkaAuthentication() *kafka.Dialer {
	precisaLogin, err := strconv.ParseBool(utils.GetVariavelAmbiente("kafka_need_authentication", "false"))
	if err != nil {
		return nil
	}

	if !(precisaLogin) {
		return nil
	}

	username := utils.GetVariavelAmbiente("kafka_username", "username")
	password := utils.GetVariavelAmbiente("kafka_password", "senha")
	return &kafka.Dialer{
		SASLMechanism: plain.Mechanism{
			Username: username,
			Password: password,
		},
	}
}
