package domain

import "time"

type BusinessMessage struct {
	ChaveAcessoCte    string     `json:"chave_acesso_cte"`
	CnpjTransportador string     `json:"cnpj_transportador"`
	TpEvento          string     `json:"tpEvento"`
	TipoErro          []TipoErro `json:"tipo_erro"`
	DhEvento          string     `json:"dhEvento"`
	NProt             string     `json:"nProt"`
	XJust             string     `json:"xJust"`
	DescEvento        string     `json:"descEvento"`
	DataRecebimento   time.Time  `json:"data_recebimento"`
}
type TipoErro struct {
	Campo string `json:"campo"`
}
