package main

import (
	"./models"
	"fmt"
)

func main () {

	//tarjeta, err := models.GetTarjetaByID(2)
	//checkError(err)
	//fmt.Println(tarjeta)
	//tarjeta.Eliminar()
	tarjetas,_ := models.GetTarjetas()
	fmt.Println(tarjetas)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}