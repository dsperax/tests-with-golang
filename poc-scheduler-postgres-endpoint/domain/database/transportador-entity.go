package domain

import "time"

type Transportador struct {
	CdEmpresa             int       `json:"cd_empresa"`
	CdCnpjTransportador   string    `json:"cd_cnpj_transportador"`
	DsEmailNotificacaoPgt string    `json:"ds_email_notificacao_pgt"`
	StatusAtivo           string    `json:"status_ativo"`
	CdMatriculaCadastro   int       `json:"cd_matricula_cadastro"`
	DhCadastro            time.Time `json:"dh_cadastro"`
	CdMatriculaManutencao int       `json:"cd_matricula_manutencao"`
	DhManutencao          time.Time `json:"dh_manutencao"`
}

func (*Transportador) TableName() string {
	return "tb_cadastro_transportador"
}
