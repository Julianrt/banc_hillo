package main

import (
    "log"
    "net/http"

    "./routes"
    
    "github.com/gorilla/mux"
)

func main () {
    mux := mux.NewRouter()
    routes.Endpoints(mux)

    assets := http.FileServer(http.Dir("assets"))
    statics := http.StripPrefix("/assets/", assets)
    mux.PathPrefix("/assets/").Handler(statics)

    log.Println("El servidor está escuchando por el puerto :8000")
    server := http.Server{
        Addr: 		":8000",
        Handler: 	mux,
    }
    log.Fatal(server.ListenAndServe())
}
