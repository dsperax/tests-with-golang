package repository

import (
	"context"
	"encoding/json"
	"fmt"

	domain "poc-consumer-go-mongo-sync/domain/database"
	"poc-consumer-go-mongo-sync/domain/log"
	"poc-consumer-go-mongo-sync/infra/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BusinessRepoMongoDb struct {
}

func NewBusinessMongoDbRepository() IBusinessRepository {
	return &BusinessRepoMongoDb{}
}

func (r *BusinessRepoMongoDb) Inserir(businessdomain domain.BusinessEntity) error {
	coll := database.GetCollection("poc_golang_teste")
	doc := businessdomain

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}
	fmt.Printf("Inserido com sucesso _id: %v\n", result.InsertedID)

	return nil
}

func (r *BusinessRepoMongoDb) Excluir(businessdomain domain.BusinessEntity) error {
	// err := database.DB.Where(
	// 	"cd_empgcb_atd=? and cd_est_sig_zen=? and cd_zen=?",
	// 	businessdomain.Empresa,
	// 	businessdomain.UFZona,
	// 	businessdomain.CodigoZona).Delete(&businessdomain).Error
	// if err != nil {
	// 	log.GravarLog(err.Error())
	// 	return err
	// }

	return nil
}

func (r *BusinessRepoMongoDb) ExisteRegistro(businessdomain domain.BusinessEntity) (bool, error) {
	coll := database.GetCollection("poc_golang_teste")
	filter := bson.D{
		{Key: "ufzona", Value: businessdomain.UFZona},
		{Key: "codigozona", Value: businessdomain.CodigoZona}}

	var result bson.M
	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("Nenhum documento encontrado %s\n", result)
		return false, nil
	}

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}

	fmt.Println("Registro encontrado!")
	fmt.Printf("%s\n", jsonData)
	return true, nil
}

func (r *BusinessRepoMongoDb) Atualizar(businessdomain domain.BusinessEntity) error {
	// err := database.DB.Where(
	// 	"cd_empgcb_atd=? and cd_est_sig_zen=? and cd_zen=?",
	// 	businessdomain.Empresa,
	// 	businessdomain.UFZona,
	// 	businessdomain.CodigoZona).Save(&businessdomain).Error
	// if err != nil {
	// 	log.GravarLog(err.Error())
	// 	return err
	// }

	return nil
}
