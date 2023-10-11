package akafka

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"poc-consumer-go-endpoint-integrador/app/utils"
	"poc-consumer-go-endpoint-integrador/domain/log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
)

func Consumer(topic string, servers []string, groupid string, mensagem chan kafka.Message, resultado chan bool) {
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
			break
		}

		mensagem <- msg

		if <-resultado {
			kafkaConsumer.CommitMessages(ctx, msg)
		} else {
			continue
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
