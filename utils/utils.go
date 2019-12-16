package utils

import (
	"net/http"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func HideCard(tarjeta string) string {
	s := strings.Split(tarjeta, "")
	hidden:="************"
	for i:=12; i<len(s); i++{
		hidden+=s[i]
	}
	return hidden
}
