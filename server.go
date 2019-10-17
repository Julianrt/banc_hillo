package main

import (
	"./models"
	"fmt"
)

func main () {
	
	cuenta,_ := models.CrearEmpleado("jkl","jkl", "jkl", "jkl","123")
	models.CrearEmpleado("asd","asd", "asd", "asd","123")
	cuenta2,_ := models.CrearEmpleado("asd","asd", "asd", "asdd","123")

	empleado, _ := models.GetEmpleadoByID(cuenta.ID)

	//fmt.Println(empleado)

	empleado.SetPassword("321")

	//fmt.Println(empleado)

	cuenta2.LogicDelete()

	fmt.Println(models.GetEmpleados())
	fmt.Println(models.GetEmpleadoByUsername("asd"))
	fmt.Println(models.GetEmpleadoByUsername("asdd"))

}

func checkError(err error) {
	fmt.Println(err)
}