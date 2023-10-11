package database

import (
	"context"
	"fmt"
	"os"

	"poc-consumer-go-mongo-sync/app/utils"
	"poc-consumer-go-mongo-sync/domain/log"

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
			"aa",
			"bb",
			"cc",
			27017,
			"nome",
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

	Banco = client.Database("nome")

	log.GravarLog("Connected to MongoDB")
}

func GetCollection(nome string) *mongo.Collection {
	return Banco.Collection(nome)
}
