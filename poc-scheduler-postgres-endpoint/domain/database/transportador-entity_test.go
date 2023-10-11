package domain

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	transportador Transportador = Transportador{
		CdCnpjTransportador:   "123",
		DsEmailNotificacaoPgt: "email@email.com",
	}
)

func TestTableTransportador(t *testing.T) {
	type testes struct {
		entrada  Transportador
		esperado string
	}

	execucoes := []testes{
		{
			entrada:  transportador,
			esperado: "tb_cadastro_transportador",
		},
		{
			entrada:  transportador,
			esperado: "tb_cadastro_transportador",
		},
	}

	for i, teste := range execucoes {
		t.Run(fmt.Sprintf("teste %d", i), func(t *testing.T) {
			transportador.TableName()
			assert.Equal(t, teste.esperado, teste.entrada.TableName())
		})
	}
}
