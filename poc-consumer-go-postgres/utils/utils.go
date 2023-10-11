package utils

import (
	"strconv"
	"strings"
	"time"
)

func ValidaInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		s = "0"
	}
	return strconv.Atoi(s)
}

func ValidaInt64(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		s = "0"
	}
	return strconv.ParseInt(s, 10, 64)
}

func ValidaData(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		s = "0001-01-01"
	}
	return time.Parse("2006-01-02", s)
}

func ValidaDataHora(s string) (time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		s = "2006-01-02T15:04:05"
	}
	return time.Parse("2006-01-02T15:04:05", s)
}
