# Projects with go

- poc-consumer-go-endpoint-integrador: Consumer que envia objeto para o integrador realizar a presistência;
- poc-consumer-go-postgres: Consome a mensagem e realiza a persistência;
- poc-consumer-go-mongo-sync: Consome mensagens uma a uma;
- poc-consumer-go-mongo-async: Consome mensagens por bloco e realiza a persistencia no final;
- poc-scheduler-postgres-endpoint: scheduler que extrai dados da base, realiza logica e chama endpoint para disparo de emails;

## Log generic class ex

```
package log

import (
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Log struct {
	TsGravacao time.Time `json:"tsLog"`
	Conteudo   string    `json:"conteudo"`
}

func GravaLogSucesso(msgSucesso string) {
	debug := false

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().
		Msg(msgSucesso)

	// Debug log
	log.Debug().Msg("Exiting Program")
}

func GravaLogErro(registro string, erro string) {
	debug := false

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Error().
		Str("erro", erro).
		Msg("Erro ao processar o registro: " + registro)

	// Debug log
	log.Debug().Msg("Exiting Program")
}```