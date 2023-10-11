package repository

import (
	"fmt"
	"testing"

	domain "poc-scheduler-postgres-endpoint/domain/database"

	"github.com/pashagolub/pgxmock"
)

func setupFetchTransportador() ([]string, domain.Transportador, pgxmock.PgxPoolIface) {
	obj := domain.Transportador{
		CdCnpjTransportador:   "123",
		DsEmailNotificacaoPgt: "email@email.com",
	}
	cols := []string{obj.CdCnpjTransportador, obj.DsEmailNotificacaoPgt}

	mock, _ := pgxmock.NewPool()

	return cols, obj, mock
}

func TestGetTransportador(t *testing.T) {
	cols, obj, mock := setupFetchTransportador()
	mock.ExpectQuery("SELECT (.+) FROM tb_cadastro_transportador").
		WillReturnRows(pgxmock.NewRows(cols).AddRow(
			obj.CdCnpjTransportador,
			obj.DsEmailNotificacaoPgt,
		))
}

func TestFetch_QueryError_Transportador(t *testing.T) {
	_, _, mock := setupFetchTransportador()
	defer mock.Close()
	mock.ExpectQuery("SELECT (.+) FROM tb_cadastro_transportador").
		WillReturnError(fmt.Errorf("ANY QUERY ERROR"))
}
