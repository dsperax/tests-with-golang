package domain

import (
	"strconv"
	"strings"
	"time"

	domainkafka "poc-consumer-go-mongo-async/domain/kafka"
)

type BusinessEntity struct {
	UFZona       string    `json:"UFZona"`
	CodigoZona   string    `json:"CodigoZona"`
	NomeZona     string    `json:"NomeZona"`
	DataCadastro time.Time `json:"DataCadastro"`
	Empresa      int       `json:"Empresa"`
}

func (p *BusinessEntity) FromDTO(mensagemkafka domainkafka.BusinessKafka) error {
	var err error = nil

	p.UFZona = mensagemkafka.UFZona
	p.CodigoZona = mensagemkafka.CodigoZona
	p.NomeZona = strings.TrimSpace(mensagemkafka.NomeZona)

	p.Empresa, err = strconv.Atoi(mensagemkafka.Empresa)
	if err != nil {
		return err
	}

	p.DataCadastro, err = time.Parse("2006-01-02", mensagemkafka.DataCadastro)
	if err != nil {
		return err
	}

	return nil
}
