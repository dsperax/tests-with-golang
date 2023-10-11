package usecase

import (
	"sort"

	domain "poc-scheduler-postgres-endpoint/domain/database"
	"poc-scheduler-postgres-endpoint/domain/log"
	domainNot "poc-scheduler-postgres-endpoint/domain/notification"
	"poc-scheduler-postgres-endpoint/infra/repository"
	"poc-scheduler-postgres-endpoint/utils"

	"golang.org/x/exp/slices"
)

type BusinessUseCase struct {
}

var (
	transportadorRepo repository.ITransportadorRepository
	cteErroRepo       repository.ICteErroRepository
	descErroRepo      repository.IDescricaoErrosRepository
	notificaRepo      repository.IBusinessRepository
)

func NewBusinessUseCase(transpRepo repository.ITransportadorRepository, cteComErroRepo repository.ICteErroRepository,
	descricaoErroRepo repository.IDescricaoErrosRepository, businessRepo repository.IBusinessRepository) INotificaConsumerUseCase {
	transportadorRepo = transpRepo
	cteErroRepo = cteComErroRepo
	descErroRepo = descricaoErroRepo
	notificaRepo = businessRepo
	return &BusinessUseCase{}
}

func (s *BusinessUseCase) Executa() {
	ctesError, err := cteErroRepo.Recuperar()
	if err != nil {
		log.GravarLog("Erro ao recuperar os CTe's com Erro: " + err.Error())
		return
	}
	if len(ctesError) == 0 {
		log.GravarLog("Não existem CTe's com status de ERRO [E]")
		return
	}

	ctesErrorParaNotificar, idCtesErrorVencidas := s.IdentificaCteVencida(ctesError)

	ctesErrorParaNotificar = s.AgrupaTransportadores(ctesErrorParaNotificar)

	listaTranportadores := s.GeraListaCnpjs(ctesErrorParaNotificar)

	err = s.NotificaTransportador(ctesErrorParaNotificar, listaTranportadores)
	if err != nil {
		log.GravarLog("Erro ao recuperar descrição dos erros: " + err.Error())
		return
	}

	err = cteErroRepo.AtualizaStatus(idCtesErrorVencidas)
	if err != nil {
		log.GravarLog("Erro ao atualizar status de cte's vencidas (mais que 7 dias): " + err.Error())
		return
	}
}

func (w *BusinessUseCase) GeraListaCnpjs(ctes []domain.CteErro) []string {
	var listaCnpjs []string
	for _, cte := range ctes {
		if !slices.Contains(listaCnpjs, cte.CnpjTransportador) {
			listaCnpjs = append(listaCnpjs, cte.CnpjTransportador)
		}
	}
	return listaCnpjs
}

func (v *BusinessUseCase) AgrupaTransportadores(ctes []domain.CteErro) []domain.CteErro {
	sort.Slice(ctes, func(i, j int) bool {
		return ctes[i].CnpjTransportador < ctes[j].CnpjTransportador
	})
	return ctes
}

func (r *BusinessUseCase) GeraNotificationCte(cte domain.CteErro, listaErros []domain.DescricaoErros) domainNot.NotificationObject {
	var notificacao domainNot.NotificationObject
	notificacao.Id = cte.IdCte
	notificacao.DtInclusao = cte.DtInclusao
	for _, erro := range listaErros {
		notificacao.Descricao = append(notificacao.Descricao, erro.DescricaoErro)
	}
	return notificacao
}

func (t *BusinessUseCase) NotificaTransportador(ctes []domain.CteErro, transportadoresDistintos []string) error {

	for _, cnpj := range transportadoresDistintos {
		var listaEnvioTransp domainNot.NotificationEntity
		transportador, err := transportadorRepo.Recuperar(cnpj)
		if err != nil {
			log.GravarLog("Erro ao recuperar descrição dos erros: " + err.Error() + " idCte: " + cnpj)
			return err
		}
		listaEnvioTransp.DsEmailNotificacaoPgt = transportador.DsEmailNotificacaoPgt
		listaEnvioTransp.CorpoEmail = "Segue abaixo a lista de CTe's aguardando correção. Vale destacar que após o periodo de 7 dias da data de recebimento, o documento NÃO poderá mais ser alterado."
		listaEnvioTransp.TituloEmail = "CTE'S COM ERRO A CORRIGIR"

		for _, cteError := range ctes {
			if cteError.CnpjTransportador == cnpj {
				descricaoErros, err := descErroRepo.Recuperar(cteError.IdCte)
				if err != nil {
					log.GravarLog("Erro ao recuperar descrição dos erros: " + err.Error() + " idCte: " + cteError.IdCte)
					return err
				}

				notificationCte := t.GeraNotificationCte(cteError, descricaoErros)
				listaEnvioTransp.ListaObjeto = append(listaEnvioTransp.ListaObjeto, notificationCte)
			}
		}

		err = notificaRepo.EndPointNotificacaoEmails(listaEnvioTransp)
		if err != nil {
			log.GravarLog("Erro ao chamar Api de Emails: " + err.Error() + " CNPJ do Transportador: " + transportador.CdCnpjTransportador)
			return err
		}
	}
	return nil
}

func (u *BusinessUseCase) IdentificaCteVencida(ctesError []domain.CteErro) ([]domain.CteErro, []domain.CteErro) {
	var ctesErrorParaNotificar []domain.CteErro
	var idCtesErrorVencidas []domain.CteErro
	for _, cte := range ctesError {
		if utils.CompareDates(cte.DtInclusao) {
			idCtesErrorVencidas = append(idCtesErrorVencidas, cte)
		} else {
			ctesErrorParaNotificar = append(ctesErrorParaNotificar, cte)
		}
	}
	return ctesErrorParaNotificar, idCtesErrorVencidas
}
