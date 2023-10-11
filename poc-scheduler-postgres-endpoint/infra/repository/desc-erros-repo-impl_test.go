package repository

import (
	"fmt"
	"testing"

	domain "poc-scheduler-postgres-endpoint/domain/database"

	"github.com/pashagolub/pgxmock"
)

func setupFetchCteErroDescricao() ([]string, domain.DescricaoErros, pgxmock.PgxPoolIface) {
	obj := domain.DescricaoErros{
		Id:            "id",
		IdCte:         "id_cte",
		DescricaoErro: "obj Teste",
	}
	cols := []string{obj.Id, obj.IdCte, obj.DescricaoErro}

	mock, _ := pgxmock.NewPool()

	return cols, obj, mock
}

func TestGetCteErroDescricao(t *testing.T) {
	cols, obj, mock := setupFetchCteErroDescricao()
	mock.ExpectQuery("SELECT (.+) FROM tb_erro_cte_descricao").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			obj.Id,
			obj.IdCte,
			obj.DescricaoErro,
		))
}

func TestFetch_QueryError_CteError_Descricao(t *testing.T) {
	_, _, mock := setupFetchCteErroDescricao()
	defer mock.Close()
	mock.ExpectQuery("SELECT (.+) FROM tb_erro_cte_descricao").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))
}
