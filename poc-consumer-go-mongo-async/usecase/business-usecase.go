package usecase

import (
	"encoding/json"
	"fmt"

	domaindatabase "poc-consumer-go-mongo-async/domain/database"
	domain "poc-consumer-go-mongo-async/domain/kafka"
	"poc-consumer-go-mongo-async/domain/log"
	"poc-consumer-go-mongo-async/infra/repository"

	"github.com/segmentio/kafka-go"
)

type BusinessUseCase struct {
}

var (
	canalresultado chan bool
	businessrepo   repository.IBusinessRepository
)

func NewBusinessUseCase(repository repository.IBusinessRepository, canalresult chan bool) IKafkaConsumerUseCase {
	canalresultado = canalresult
	businessrepo = repository
	return &BusinessUseCase{}
}

func (s *BusinessUseCase) Executa(mensagemkafka kafka.Message) {
	log.GravarLog("Processo executa do usecase")
	log.GravarLog(fmt.Sprintf("Mensagem: %s", string(mensagemkafka.Value)))

	businesskafkadomain, businessentitydomain, err := s.geraEntidade(mensagemkafka)
	if err != nil {
		log.GravarLog(err.Error())
		canalresultado <- true
		return
	}

	if businesskafkadomain.Empresa != "21" {
		canalresultado <- true
		return
	}

	if businesskafkadomain.Operacao == "D" {
		err = businessrepo.Excluir(businessentitydomain)
	} else {
		log.GravarLog("Validando se o registro existe")
		existe, err := businessrepo.ExisteRegistro(businessentitydomain)
		if err != nil {
			canalresultado <- false
			return
		}

		if existe {
			log.GravarLog("Tentando atualizar o registro")
			err = businessrepo.Atualizar(businessentitydomain)
		} else {
			log.GravarLog("Tentando incluir o registro")
			err = businessrepo.Inserir(businessentitydomain)
		}
	}

	log.GravarLog(fmt.Sprintf("Mensagem processada: Op[%s] %d %s %s",
		businesskafkadomain.Operacao,
		businessentitydomain.Empresa,
		businessentitydomain.UFZona,
		businessentitydomain.CodigoZona))
	canalresultado <- err == nil
}

func (r *BusinessUseCase) geraEntidade(mensagemkafka kafka.Message) (domain.BusinessKafka, domaindatabase.BusinessEntity, error) {
	var businessentitydomain domaindatabase.BusinessEntity
	var businesskafkadomain domain.BusinessKafka
	err := json.Unmarshal(mensagemkafka.Value, &businesskafkadomain)
	if err != nil {
		log.GravarLog(err.Error())
		return businesskafkadomain, businessentitydomain, err
	}

	err = businessentitydomain.FromDTO(businesskafkadomain)
	if err != nil {
		log.GravarLog(err.Error())
		return businesskafkadomain, businessentitydomain, err
	}

	return businesskafkadomain, businessentitydomain, nil
}
