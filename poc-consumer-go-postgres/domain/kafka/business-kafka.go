package domain

type BusinessKafka struct {
	DocumentoOrigem    string `json:"CD_TRKUNI_DOC_ORI"`
	CodigoDocumento    string `json:"CD_TRKUNI_DOC"`
	Desmembramento     string `json:"CD_TRKUNI_DOC_SEQ"`
	NumeroEntrega      string `json:"CD_TRKUNI_NUM_ETG"`
	PontoControle      string `json:"CD_TRKUNI_PTO_CTE"`
	TsOcorrencia       string `json:"TS_TRKUNI_AUX"`
	CnpjTransportadora string `json:"CD_TRKUNI_CGC_TRP"`
	NumeroConhecimento string `json:"CD_TRKUNI_NUM_CNM"`
	NumeroNota         string `json:"CD_TRKUNI_NUM_NFI"`
	SerieNota          string `json:"CD_TRKUNI_SER_NFI"`
	DataNota           string `json:"DT_TRKUNI_EMI_NFI"`
	ChaveAcesso        string `json:"CD_TRKUNI_CHA_NFI"`
	CnpjEmitente       string `json:"CD_TRKUNI_CGC_EMI"`
	EmpresaEmitente    string `json:"CD_EMPGCB_EMI"`
	FilialEmitente     string `json:"CD_FIL_EMI"`
	AnoEmissao         string `json:"AA_TRKUNI_EMI"`
	MesEmissao         string `json:"MM_TRKUNI_EMI"`
	PagamentoFrete     string `json:"ST_TRKUNI_PGT_FRT"`
	PontoFinalizador   string `json:"ST_TRKUNI_TRK_FIN"`
	Operacao           string `json:"AUD_ENTTYP"`
}
