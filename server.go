package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cuenta, _ := models.AltaCuenta(9696, 1234, "Fulanito", 1)

	usuario, _ := models.CrearUsuario("Mangano", "Manganito", "123")

	fmt.Println(cuenta)
	fmt.Println(usuario)

}

func checkError(err error) {
	fmt.Println(err)
}