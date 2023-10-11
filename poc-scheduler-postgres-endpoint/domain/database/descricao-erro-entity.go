package domain

type DescricaoErros struct {
	Id            string `json:"id"`
	IdCte         string `json:"id_cte"`
	DescricaoErro string `json:"descricao_erro"`
}

func (*DescricaoErros) TableName() string {
	return "tb_erro_cte_descricao"
}
