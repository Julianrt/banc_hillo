package routers

import (
	"../handlers"
	"github.com/gorilla/mux"
)

func Endpoints(mux *mux.Router) {
	clientesEndpoints(mux)
	tiposCuentaEndpoints(mux)
	cuentasEndpoints(mux)
	tarjetasEndpoints(mux)
	empleadosEndpoints(mux)
	tiposTransaccionesEndpoints(mux)
	transaccionEndpoints(mux)
}

func clientesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/clientes/", handlers.GetClientes).Methods("GET")
	mux.HandleFunc("/api/clientes/{id:[0-9]+}", handlers.GetCliente).Methods("GET")
	mux.HandleFunc("/api/clientes/", handlers.CreateCliente).Methods("POST")
}

func cuentasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/cuentas/", handlers.GetCuentas).Methods("GET")
	mux.HandleFunc("/api/cuentas/{id:[0-9]+}", handlers.GetCuenta).Methods("GET")
	mux.HandleFunc("/api/cuentas/", handlers.CreateCuenta).Methods("POST")
}

func tarjetasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tarjetas/", handlers.GetTarjetas).Methods("GET")
	mux.HandleFunc("/api/tarjetas/{id:[0-9]+}", handlers.GetTarjeta).Methods("GET")
	mux.HandleFunc("/api/tarjetas/", handlers.CreateTarjeta).Methods("POST")
}

func tiposCuentaEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_cuenta/", handlers.GetTiposCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/{id:[0-9]+}", handlers.GetTipoCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/", handlers.CreateTipoCuenta).Methods("POST")
}

func tiposTransaccionesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_transaccion/", handlers.GetTiposTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/{id:[0-9]+}", handlers.GetTipoTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/", handlers.CreateTipoTransaccion).Methods("POST")
}

func empleadosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/empleados/", handlers.GetEmpleados).Methods("GET")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.GetStatus).Methods("GET")
	//mux.HandleFunc("/api/v1/status/", handlers.CreateCuenta).Methods("POST")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.UpdateStatus).Methods("PUT")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.DeleteStatus).Methods("DELETE")
}

func transaccionEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/transacciones/", handlers.GetTransacciones).Methods("GET")
	mux.HandleFunc("/api/transacciones/transferencias/", handlers.DoTransferencia).Methods("POST")
	mux.HandleFunc("/api/transacciones/depositos/", handlers.DoDeposito).Methods("POST")
}
