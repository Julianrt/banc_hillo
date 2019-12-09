package models

import (
	"time"
	"fmt"
)

func ObtenerFechaHoraActualString() string {
	t := time.Now()

	fechaHora := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", 
		t.Year(), t.Month(), t.Day(), 
		t.Hour(), t.Minute(), t.Second())

	return fechaHora
}

func GetFechaVencimientoString() string {
	t := time.Now()
	mes := fmt.Sprintf("%02d",t.Month())
	year := fmt.Sprintf("%02d",t.Year()-1998)
	fecha := mes+"/"+year
	return fecha
}
