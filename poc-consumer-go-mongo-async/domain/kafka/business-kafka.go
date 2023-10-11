package domain

type BusinessKafka struct {
	Operacao     string `json:"AUD_ENTTYP"`
	DataHora     string `json:"AUD_APPLY_TIMESTAMP"`
	UFZona       string `json:"CD_EST_SIG_ZEN"`
	CodigoZona   string `json:"CD_ZEN"`
	NomeZona     string `json:"NM_ZEN"`
	DataCadastro string `json:"DT_ZEN_CAD"`
	Empresa      string `json:"CD_EMPGCB_ATD"`
    Objeto       string `json:"OBJETO"`
	Entidade     string `json:"ENTTYP"`
}
