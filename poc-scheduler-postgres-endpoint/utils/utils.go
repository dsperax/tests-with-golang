package utils

import (
	"time"

	"poc-scheduler-postgres-endpoint/domain/log"
)

func CompareDates(dateString string) bool {
	layout := "2006-01-02T15:04:05Z"
	today := time.Now()

	date, err := time.Parse(layout, dateString)
	if err != nil {
		log.GravarLog("Erro ao fazer parsing da data: " + err.Error())
		return false
	}

	difference := today.Sub(date)
	sevenDays := 7 * 24 * time.Hour

	return difference > sevenDays
}
