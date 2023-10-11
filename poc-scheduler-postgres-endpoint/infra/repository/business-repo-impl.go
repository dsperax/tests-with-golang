package repository

import (
	"bytes"
	"encoding/json"
	"net/http"

	"poc-scheduler-postgres-endpoint/app/utils"
	"poc-scheduler-postgres-endpoint/domain/log"
	domain "poc-scheduler-postgres-endpoint/domain/notification"
)

const ()

type BusinessRepoHttp struct {
}

func NewBusinessRepository() IBusinessRepository {
	return &BusinessRepoHttp{}
}

func (r *BusinessRepoHttp) EndPointNotificacaoEmails(notificacao domain.NotificationEntity) error {
	httpposturl := utils.GetVariavelAmbiente("endpoint", "http://localhost:8083/disparo-email")

	jsonData, err := json.Marshal(notificacao)
	if err != nil {
		log.GravarLog("Erro Marshal endpoint: " + err.Error())
		return err
	}

	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.GravarLog("Erro na chamada do EndPointRegras: " + err.Error())
		return err
	}
	defer response.Body.Close()

	log.GravarLog("Response Status: " + response.Status)
	return nil
}
