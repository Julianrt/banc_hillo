package routers

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func Endpoints(mux *mux.Router) {
	empleadosEndpoints(mux)
	transaccionEndpoints(mux)
}

func empleadosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/empleados/", handlers.GetEmpleados).Methods("GET")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.GetStatus).Methods("GET")
	//mux.HandleFunc("/api/v1/status/", handlers.CreateCuenta).Methods("POST")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.UpdateStatus).Methods("PUT")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.DeleteStatus).Methods("DELETE")
}

func transaccionEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/transferencia/", handlers.DoTransferencia).Methods("POST")
}
