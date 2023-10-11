package repository

import (
	"fmt"

	domain "poc-consumer-go-postgres/domain/database"
	"poc-consumer-go-postgres/domain/log"
	"poc-consumer-go-postgres/infra/database"
)

const (
	filtrochave = "DOCUMENTO_ORIGEM=? and CODIGO_DOCUMENTO=? and DESMEMBRAMENTO=? and PONTO_CONTROLE=? and TS_OCORRENCIA=?"
)

type BusinessRepoPg struct {
}

func NewBusinessMSSQLRepository() IBusinessRepository {
	return &BusinessRepoPg{}
}

func (r *BusinessRepoPg) Inserir(businessdomain domain.BusinessEntity) error {
	err := database.DB.Create(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}

func (r *BusinessRepoPg) Excluir(businessdomain domain.BusinessEntity) error {
	err := database.DB.Where(
		filtrochave,
		businessdomain.DocumentoOrigem,
		businessdomain.CodigoDocumento,
		businessdomain.Desmembramento,
		businessdomain.PontoControle,
		businessdomain.TsOcorrencia).Delete(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}

func (r *BusinessRepoPg) ExisteRegistro(businessdomain domain.BusinessEntity) (bool, error) {
	err := database.DB.Where(
		filtrochave,
		businessdomain.DocumentoOrigem,
		businessdomain.CodigoDocumento,
		businessdomain.Desmembramento,
		businessdomain.PontoControle,
		businessdomain.TsOcorrencia).First(&businessdomain).Error
	if err != nil && err.Error() != "record not found" {
		log.GravarLog(fmt.Sprintf("Deu erro no find: %s", err.Error()))
		return false, err
	}

	if err != nil && err.Error() == "record not found" {
		return false, nil
	}

	return true, nil
}

func (r *BusinessRepoPg) Atualizar(businessdomain domain.BusinessEntity) error {
	err := database.DB.Where(
		filtrochave,
		businessdomain.DocumentoOrigem,
		businessdomain.CodigoDocumento,
		businessdomain.Desmembramento,
		businessdomain.PontoControle,
		businessdomain.TsOcorrencia).Save(&businessdomain).Error
	if err != nil {
		log.GravarLog(err.Error())
		return err
	}

	return nil
}
