package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cuenta1, err := models.GetCuentaByNumeroCuenta(1234)
	checkError(err)
	fmt.Println(cuenta1)

	cuenta2, err := models.GetCuentaByNumeroCuenta(4444)
	checkError(err)
	fmt.Println(cuenta2)

	/*err = cuenta1.Transferir(cuenta2.NumeroDeCuenta, 56.4)
	checkError(err)

	fmt.Println(cuenta1)
*/
	fmt.Println(cuenta1.SolicitarSaldo())

}

func checkError(err error) {
	fmt.Println(err)
}