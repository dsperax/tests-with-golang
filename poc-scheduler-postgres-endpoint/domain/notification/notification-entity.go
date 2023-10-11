package domain

type NotificationEntity struct {
	DsEmailNotificacaoPgt string               `json:"email_notificacao"`
	TituloEmail           string               `json:"titulo_email"`
	CorpoEmail            string               `json:"corpo_email"`
	ListaObjeto           []NotificationObject `json:"lista_objeto"`
}
