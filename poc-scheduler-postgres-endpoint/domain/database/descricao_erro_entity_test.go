package domain

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	erro DescricaoErros = DescricaoErros{
		Id:            "id",
		IdCte:         "id_cte",
		DescricaoErro: "Erro Teste",
	}
)

func TestTableDescricaoErros(t *testing.T) {
	type testes struct {
		entrada  DescricaoErros
		esperado string
	}

	execucoes := []testes{
		{
			entrada:  erro,
			esperado: "tb_erro_cte_descricao",
		},
		{
			entrada:  erro,
			esperado: "tb_erro_cte_descricao",
		},
	}

	for i, teste := range execucoes {
		t.Run(fmt.Sprintf("teste %d", i), func(t *testing.T) {
			erro.TableName()
			assert.Equal(t, teste.esperado, teste.entrada.TableName())
		})
	}
}
