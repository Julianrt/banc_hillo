package main

import (
	"log"
	"net/http"
	"./routers"
	"github.com/gorilla/mux"
	
	"./models"
)

func main () {
	models.CrearCliente("Cine","","","RFC")
	models.AltaCuenta("1234",1,2)
	models.CrearTarjeta(1,1,"0000", "1111", "11/21", "123")
	cuenta1,_ := models.GetCuentaByID(1)
	cuenta1.Depositar(10590)

	models.CrearCliente("pepito","","","CURP")
	models.AltaCuenta("4444",2,1)
	models.CrearTarjeta(2,2,"9999", "1111", "11/21", "123")
	cuenta2,_ := models.GetCuentaByID(2)
	cuenta2.Depositar(3600)




	mux := mux.NewRouter()
	routers.Endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :8000")
	server := http.Server{
		Addr: 		":8000",
		Handler: 	mux,
	}
	log.Fatal(server.ListenAndServe())
}
