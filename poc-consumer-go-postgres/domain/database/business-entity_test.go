package domain

import (
	"fmt"
	"testing"

	domainkafka "poc-consumer-go-postgres/domain/kafka"

	"github.com/stretchr/testify/assert"
)

var (
	kfkMsg domainkafka.BusinessKafka = domainkafka.BusinessKafka{
		DocumentoOrigem:    "123",
		CodigoDocumento:    "123",
		Desmembramento:     "123",
		NumeroEntrega:      "123",
		PontoControle:      "123",
		TsOcorrencia:       "2023-05-21T01:01:12.694615000000",
		CnpjTransportadora: "123",
		NumeroConhecimento: "123",
		NumeroNota:         "123",
		SerieNota:          "123",
		DataNota:           "",
		ChaveAcesso:        "123",
		CnpjEmitente:       "123",
		EmpresaEmitente:    "123",
		FilialEmitente:     "123",
		AnoEmissao:         "123",
		MesEmissao:         "123",
		PagamentoFrete:     "123",
		PontoFinalizador:   "123",
		Operacao:           "123",
	}
)

func TestKafkaDatabase(t *testing.T) {
	entityObj := BusinessEntity{}
	entityObj.FromDTO(kfkMsg)
	type testes struct {
		entrada   domainkafka.BusinessKafka
		resultado BusinessEntity
		esperado  int
	}

	execucoes := []testes{
		{
			entrada:   kfkMsg,
			resultado: BusinessEntity{},
			esperado:  123,
		},
		{
			entrada:   kfkMsg,
			resultado: BusinessEntity{},
			esperado:  1234,
		},
	}

	for i, teste := range execucoes {
		t.Run(fmt.Sprintf("teste %d", i), func(t *testing.T) {
			teste.resultado.FromDTO(teste.entrada)
			assert.Equal(t, teste.esperado, teste.resultado.CodigoDocumento, "Teste Codigo Documento")
		})
	}
}
