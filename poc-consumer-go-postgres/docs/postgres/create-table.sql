create table TrackingNotaFiscal(
	DOCUMENTO_ORIGEM VARCHAR(3) not null,
	CODIGO_DOCUMENTO INT not null,
	DESMEMBRAMENTO INT not null,
	NUMERO_ENTREGA BIGINT not null,
	PONTO_CONTROLE VARCHAR(3) not null,
	TS_OCORRENCIA TIMESTAMP not null,
	CNPJ_TRANSPORTADORA VARCHAR(14),
	NUMERO_CONHECIMENTO INT,
	NUMERO_NOTA INT,
	SERIE_NOTA VARCHAR(3),
	DATA_NOTA TIMESTAMP,
	CHAVE_ACESSO VARCHAR(44),
	CNPJ_EMITENTE VARCHAR(14),
	EMPRESA_EMITENTE INT,
	FILIAL_EMITENTE INT,
	ANO_EMISSAO INT,
	MES_EMISSAO INT,
	PAGAMENTO_FRETE VARCHAR(1),
	PONTO_FINALIZADOR VARCHAR(1))