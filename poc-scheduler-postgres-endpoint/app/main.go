package main

import (
	"time"

	"poc-scheduler-postgres-endpoint/app/router"
	"poc-scheduler-postgres-endpoint/domain/log"
	"poc-scheduler-postgres-endpoint/infra/database"
	"poc-scheduler-postgres-endpoint/infra/repository"
	"poc-scheduler-postgres-endpoint/usecase"

	"github.com/go-co-op/gocron"
)

var (
	transportadorRepo repository.ITransportadorRepository  = repository.NewTransportadorPgRepository()
	cteErroRepo       repository.ICteErroRepository        = repository.NewCteErroPgRepository()
	descErroRepo      repository.IDescricaoErrosRepository = repository.NewDescricaoErrosPgRepository()
	businessrepo      repository.IBusinessRepository       = repository.NewBusinessRepository()
	service           usecase.INotificaConsumerUseCase     = usecase.NewBusinessUseCase(transportadorRepo, cteErroRepo, descErroRepo, businessrepo)
)

func main() {
	go router.Roteador()
	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(1).Day().At("13:00").Do(func() {
		// _, _ = s.Every(60).Seconds().Do(func() {
		log.GravarLog("Execução da rotina de notificação de transportadores para CTe's com erro.")
		database.ConectaComBancoDeDados()
		service.Executa()
	})
	s.StartBlocking()
}
