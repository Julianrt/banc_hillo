package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cliente, _ := models.CrearCliente("Julián","Ruiz","Tequida", "RUTJ960918HSRZQL01")

	nuevaCuenta,_ := models.AltaCuenta("0000",cliente.ID,1)

	fmt.Println(nuevaCuenta)
	
	nuevaCuenta.Depositar(6000.50)
	nuevaCuenta.Depositar(9500.00)

	fmt.Println("+++++++++++++")
	fmt.Println(cliente)
	fmt.Println(nuevaCuenta)


}

func checkError(err error) {
	fmt.Println(err)
}