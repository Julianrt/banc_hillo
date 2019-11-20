package main

import (
	"log"
	"net/http"
	"./routers"
	"github.com/gorilla/mux"
)

func main () {
	mux := mux.NewRouter()
	routers.Endpoints(mux)

	log.Println("El servidor está escuchando por el puerto :8000")
	server := http.Server{
		Addr: 		":8000",
		Handler: 	mux,
	}
	log.Fatal(server.ListenAndServe())
}
