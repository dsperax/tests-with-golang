package database

import (
	"context"
	"fmt"
	"os"

	"poc-consumer-go-mongo-async/app/utils"
	"poc-consumer-go-mongo-async/domain/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Banco *mongo.Database
	ctx   = context.TODO()
)

func ConectaComBancoDeDados() {
	stringDeConexao := utils.GetVariavelAmbiente("MONGODB_URI",
		fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=%s",
			"123",
			"123",
			"123",
			27017,
			"nomebanco",
			"admin"))

	clientOptions := options.Client().ApplyURI(stringDeConexao)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.GravarLog(err)
		os.Exit(1)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.GravarLog(err)
		os.Exit(1)
	}

	Banco = client.Database("nomebvanco")

	log.GravarLog("Connected to MongoDB")
}

func GetCollection(nome string) *mongo.Collection {
	return Banco.Collection(nome)
}
