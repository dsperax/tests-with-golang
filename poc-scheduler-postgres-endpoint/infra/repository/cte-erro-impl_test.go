package repository

import (
	"fmt"
	"testing"

	domain "poc-scheduler-postgres-endpoint/domain/database"

	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
)

func setupFetchCteErro() ([]string, domain.CteErro, pgxmock.PgxPoolIface) {
	obj := domain.CteErro{
		IdCte:                 "id_cte",
		ChaveAcessoNotaFiscal: "chave_acesso_nota_fiscal",
		CnpjTransportador:     "cnpj_transportador",
		CdErro:                "cd_erro",
		Status:                "status",
		DtInclusao:            "2023-07-21",
		DtAlteracao:           "2023-05-21",
	}
	cols := []string{obj.IdCte, obj.ChaveAcessoNotaFiscal, obj.CnpjTransportador, obj.CdErro, obj.Status, obj.DtInclusao, obj.DtAlteracao}
	fakeTransportadorDBResponse := domain.Transportador{}
	faker.FakeData(&fakeTransportadorDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, obj, mock
}

func TestGetCteErro(t *testing.T) {
	cols, obj, mock := setupFetchCteErro()
	mock.ExpectQuery("SELECT (.+) FROM tb_cte_erro").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			obj.IdCte,
			obj.ChaveAcessoNotaFiscal,
			obj.CnpjTransportador,
			obj.CdErro,
			obj.Status,
			obj.DtInclusao,
			obj.DtAlteracao,
		))
}

func TestFetch_QueryError_CteErro(t *testing.T) {
	_, _, mock := setupFetchCteErro()
	defer mock.Close()
	mock.ExpectQuery("SELECT (.+) FROM tb_cte_erro").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))
}
