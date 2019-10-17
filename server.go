package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cuenta, _ := models.CrearEmpleado("jkl","jkl", "jkl", "jkl","123")
	models.CrearEmpleado("asd","asd", "asd", "asd","123")
	models.CrearEmpleado("asd","asd", "asd", "asdd","123")

	empleado, _ := models.GetEmpleadoByID(cuenta.ID)

	//fmt.Println(empleado)

	empleado.SetPassword("321")

	//fmt.Println(empleado)

	fmt.Println(models.GetEmpleados())

}

func checkError(err error) {
	fmt.Println(err)
}