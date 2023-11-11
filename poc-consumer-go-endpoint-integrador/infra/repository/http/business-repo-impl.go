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

	data, err := removeEmptyFields(jsonData)
	if err != nil {
		log.GravarLog("Erro a remover campos vazios do JSON! " + err.Error())
		return err
	}
	
	request, _ := http.NewRequest("POST", httpposturl, bytes.NewBuffer(data))
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

func removeEmptyFields(data []byte) ([]byte, error) {
	var obj interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		return nil, err
	}

	cleanedData := removeEmpty(obj)

	result, err := json.Marshal(cleanedData)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func removeEmpty(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		cleaned := make(map[string]interface{})
		for key, value := range v {
			if value != nil {
				switch value := value.(type) {
				case string:
					if value != "" {
						cleaned[key] = removeEmpty(value)
					}
				case map[string]interface{}:
					cleanedValue := removeEmpty(value)
					if len(cleanedValue.(map[string]interface{})) > 0 {
						cleaned[key] = cleanedValue
					}
				default:
					cleaned[key] = removeEmpty(value)
				}
			}
		}
		return cleaned
	case []interface{}:
		cleaned := make([]interface{}, 0)
		for _, value := range v {
			if value != nil {
				switch value := value.(type) {
				case string:
					if value != "" {
						cleaned = append(cleaned, removeEmpty(value))
					}
				case map[string]interface{}:
					cleanedValue := removeEmpty(value)
					if len(cleanedValue.(map[string]interface{})) > 0 {
						cleaned = append(cleaned, cleanedValue)
					}
				default:
					cleaned = append(cleaned, removeEmpty(value))
				}
			}
		}
		return cleaned
	default:
		return v
	}
}
