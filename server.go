package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cuenta, _ := models.AltaCuenta(9696, 1234, "Julian Ruis", 1)

	usuario, _ := models.CrearUsuario("Julian Ruiz", "julianrt", "Carritos1")

	fmt.Println(cuenta)
	fmt.Println(usuario)

}

func checkError(err error) {
	fmt.Println(err)
}