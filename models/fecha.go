package models

import (
	"time"
	"fmt"
)

func ObtenerFechaHoraActualString() string {
	t := time.Now()

	fechaHora := fmt.Sprintf("%d-%02d-%02d - %02d:%02d:%02d", 
		t.Year(), t.Month(), t.Day(), 
		t.Hour(), t.Minute(), t.Second())

	return fechaHora
}