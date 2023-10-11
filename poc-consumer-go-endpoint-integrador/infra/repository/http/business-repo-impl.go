package repository

import (
	"bytes"
	"encoding/json"
	"net/http"

	"poc-consumer-go-endpoint-integrador/app/utils"
	"poc-consumer-go-endpoint-integrador/domain/log"
	domain "poc-consumer-go-endpoint-integrador/domain/message"
)

const ()

type BusinessRepoHttp struct {
}

func NewBusinessMSSQLRepository() IBusinessRepository {
	return &BusinessRepoHttp{}
}

func (r *BusinessRepoHttp) EndPointRegras(obj domain.BusinessMessage) error {
	httpposturl := utils.GetVariavelAmbiente("endpoint", "http://localhost:8083/evento-erro")

	jsonData, err := json.Marshal(obj)
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
