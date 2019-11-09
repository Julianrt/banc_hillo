package main

import (
	"./models"
	"fmt"
)

func main () {

	models.CrearTipoTransaccion("TRANSFERENCIA")
	models.CrearTipoTransaccion("DEPOSITO")
	
	tipoTransaccion1,_ := models.GetTipoTransaccionByID(1)
	tipoTransaccion2,_ := models.GetTipoTransaccionByID(2)

	fmt.Println(tipoTransaccion1)
	fmt.Println(tipoTransaccion2)

	tipoTransaccion2.NombreTipoTransaccion = "DEPÓSITO"
	tipoTransaccion2.Guardar()

	fmt.Println(tipoTransaccion1)
	fmt.Println(tipoTransaccion2)

	tipoTransaccion2.Eliminar()
	
	tiposTransacciones,err := models.GetTiposTransaccion()
	checkError(err)
	fmt.Println(tiposTransacciones)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}