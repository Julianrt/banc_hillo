package main

import (
	"./models"
	"fmt"
)

func main () {

	models.CrearTipoCuenta("fisica3")
	models.CrearTipoCuenta("fisica4")
	models.CrearTipoCuenta("moral3")
	models.CrearTipoCuenta("moral4")

	tipoCuena,_ := models.GetTiposCuenta()

	fmt.Println(tipoCuena)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}