package domain

import (
	"fmt"
	"time"

	domainkafka "poc-consumer-go-postgres/domain/kafka"
	"poc-consumer-go-postgres/utils"
)

type BusinessEntity struct {
	DocumentoOrigem    string    `gorm:"column:documento_origem"`
	CodigoDocumento    int       `gorm:"column:codigo_documento"`
	Desmembramento     int       `gorm:"column:desmembramento"`
	NumeroEntrega      int64     `gorm:"column:numero_entrega"`
	PontoControle      string    `gorm:"column:ponto_controle"`
	TsOcorrencia       time.Time `gorm:"column:ts_ocorrencia"`
	CnpjTransportadora string    `gorm:"column:cnpj_transportadora"`
	NumeroConhecimento int       `gorm:"column:numero_conhecimento"`
	NumeroNota         int       `gorm:"column:numero_nota"`
	SerieNota          string    `gorm:"column:serie_nota"`
	DataNota           time.Time `gorm:"column:data_nota"`
	ChaveAcesso        string    `gorm:"column:chave_acesso"`
	CnpjEmitente       string    `gorm:"column:cnpj_emitente"`
	EmpresaEmitente    int       `gorm:"column:empresa_emitente"`
	FilialEmitente     int       `gorm:"column:filial_emitente"`
	AnoEmissao         int       `gorm:"column:ano_emissao"`
	MesEmissao         int       `gorm:"column:mes_emissao"`
	PagamentoFrete     string    `gorm:"column:pagamento_frete"`
	PontoFinalizador   string    `gorm:"column:ponto_finalizador"`
}

func (*BusinessEntity) TableName() string {
	return "trackingnotafiscal"
}

func (p *BusinessEntity) FromDTO(mensagemkafka domainkafka.BusinessKafka) error {
	var err error = nil

	p.DocumentoOrigem = mensagemkafka.DocumentoOrigem
	p.CodigoDocumento, err = utils.ValidaInt(mensagemkafka.CodigoDocumento)
	if err != nil {
		return err
	}

	p.Desmembramento, err = utils.ValidaInt(mensagemkafka.Desmembramento)
	if err != nil {
		return err
	}

	p.NumeroEntrega, err = utils.ValidaInt64(mensagemkafka.NumeroEntrega)
	if err != nil {
		return err
	}

	p.PontoControle = mensagemkafka.PontoControle
	p.TsOcorrencia, err = utils.ValidaDataHora(mensagemkafka.TsOcorrencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	p.CnpjTransportadora = mensagemkafka.CnpjTransportadora
	p.NumeroConhecimento, err = utils.ValidaInt(mensagemkafka.NumeroConhecimento)
	if err != nil {
		return err
	}

	p.NumeroNota, err = utils.ValidaInt(mensagemkafka.NumeroNota)
	if err != nil {
		return err
	}

	p.SerieNota = mensagemkafka.SerieNota
	p.DataNota, err = utils.ValidaData(mensagemkafka.DataNota)
	if err != nil {
		return err
	}

	p.ChaveAcesso = mensagemkafka.ChaveAcesso
	p.CnpjEmitente = mensagemkafka.CnpjEmitente
	p.EmpresaEmitente, err = utils.ValidaInt(mensagemkafka.EmpresaEmitente)
	if err != nil {
		return err
	}

	p.FilialEmitente, err = utils.ValidaInt(mensagemkafka.FilialEmitente)
	if err != nil {
		return err
	}

	p.AnoEmissao, err = utils.ValidaInt(mensagemkafka.AnoEmissao)
	if err != nil {
		return err
	}

	p.MesEmissao, err = utils.ValidaInt(mensagemkafka.MesEmissao)
	if err != nil {
		return err
	}

	p.PagamentoFrete = mensagemkafka.PagamentoFrete
	p.PontoFinalizador = mensagemkafka.PontoFinalizador

	return nil
}
