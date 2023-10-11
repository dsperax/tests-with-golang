package usecase

import (
	"encoding/json"
	"fmt"

	"poc-consumer-go-endpoint-integrador/domain/log"
	domain "poc-consumer-go-endpoint-integrador/domain/message"
	repository "poc-consumer-go-endpoint-integrador/infra/repository/http"

	"github.com/segmentio/kafka-go"
)

type BusinessUseCase struct {
}

var (
	canalresultado chan bool
	businessrepo   repository.IBusinessRepository
)

func NewBusinessUseCase(repository repository.IBusinessRepository, canalresult chan bool) IMessagerConsumerUseCase {
	canalresultado = canalresult
	businessrepo = repository
	return &BusinessUseCase{}
}

func (s *BusinessUseCase) Executa(messageMessage kafka.Message) {
	BusinessMessagedomain, err := s.geraEntidade(messageMessage)
	if err != nil {
		log.GravarLog("Erro no metodo Executa: " + err.Error())
		canalresultado <- true
		return
	}

	err = businessrepo.EndPointRegras(BusinessMessagedomain)
	log.GravarLog(fmt.Sprintf("Mensagem processada: ID ChaveCTe[%s]",
		BusinessMessagedomain.ChaveAcessoCte))
	canalresultado <- err == nil
}

func (r *BusinessUseCase) geraEntidade(messageMessage kafka.Message) (domain.BusinessMessage, error) {
	var BusinessMessagedomain domain.BusinessMessage
	err := json.Unmarshal(messageMessage.Value, &BusinessMessagedomain)
	if err != nil {
		log.GravarLog("Erro no método geraEntidade: " + err.Error())
		return BusinessMessagedomain, err
	}
	return BusinessMessagedomain, nil
}
