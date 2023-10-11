package domain

type NotificationObject struct {
	Id         string   `json:"id"`
	DtInclusao string   `json:"dt_inclusao"`
	Descricao  []string `json:"descricao"`
}
