package domain

type CteErro struct {
	IdCte                 string `json:"id_cte"`
	ChaveAcessoNotaFiscal string `json:"chave_acesso_nota_fiscal"`
	CnpjTransportador     string `json:"cnpj_transportador"`
	CdErro                string `json:"cd_erro"`
	Status                string `json:"status"`
	DtInclusao            string `json:"dt_inclusao"`
	DtAlteracao           string `json:"dt_alteracao"`
}

func (*CteErro) TableName() string {
	return "tb_cte_erro"
}
