package domain

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
)

var (
	cteComErro CteErro = CteErro{
		IdCte:                 "id_cte",
		ChaveAcessoNotaFiscal: "chave_acesso_nota_fiscal",
		CnpjTransportador:     "cnpj_transportador",
		CdErro:                "cd_erro",
		Status:                "status",
		DtInclusao:            "2023-07-21",
		DtAlteracao:           "2023-05-21",
	}
)

func TestTableNameCteErro(t *testing.T) {
	type testes struct {
		entrada  CteErro
		esperado string
	}

	execucoes := []testes{
		{
			entrada:  cteComErro,
			esperado: "tb_cte_erro",
		},
		{
			entrada:  cteComErro,
			esperado: "tb_cte_erro",
		},
	}

	for i, teste := range execucoes {
		t.Run(fmt.Sprintf("teste %d", i), func(t *testing.T) {
			cteComErro.TableName()
			assert.Equal(t, teste.esperado, teste.entrada.TableName())
		})
	}
}
