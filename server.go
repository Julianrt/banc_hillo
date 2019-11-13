package main

import (
	"./models"
	"fmt"
)

func main () {

	login, err := models.GetLoginByToken("b1215b66da95f915b2a1c2447ef6fbe0502992dd")
	checkError(err)
	fmt.Println(login)

	login.Eliminar()

	login, err = models.GetLoginByToken("b1215b66da95f915b2a1c2447ef6fbe0502992dd")
	checkError(err)
	fmt.Println(login)

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}