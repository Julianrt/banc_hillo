package routes

import (
	"../handlers/api"
	"../handlers/app"
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

	application(mux)
}

func clientesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/clientes/", api.GetClientes).Methods("GET")
	mux.HandleFunc("/api/clientes/{id:[0-9]+}", api.GetCliente).Methods("GET")
	mux.HandleFunc("/api/clientes/", api.CreateCliente).Methods("POST")
}

func cuentasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/cuentas/", api.GetCuentas).Methods("GET")
	mux.HandleFunc("/api/cuentas/{id:[0-9]+}", api.GetCuenta).Methods("GET")
	mux.HandleFunc("/api/cuentas/", api.CreateCuenta).Methods("POST")
	mux.HandleFunc("/api/cuentas/{id:[0-9]+}", api.UpdateCuenta).Methods("PUT")
}

func tarjetasEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tarjetas/", api.GetTarjetas).Methods("GET")
	mux.HandleFunc("/api/tarjetas/{id:[0-9]+}", api.GetTarjeta).Methods("GET")
	mux.HandleFunc("/api/tarjetas/", api.CreateTarjeta).Methods("POST")
	mux.HandleFunc("/api/tarjetas/{id:[0-9]+}", api.UpdateTarjeta).Methods("PUT")
}

func tiposCuentaEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_cuenta/", api.GetTiposCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/{id:[0-9]+}", api.GetTipoCuenta).Methods("GET")
	mux.HandleFunc("/api/tipos_cuenta/", api.CreateTipoCuenta).Methods("POST")
}

func tiposTransaccionesEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/tipos_transaccion/", api.GetTiposTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/{id:[0-9]+}", api.GetTipoTransaccion).Methods("GET")
	mux.HandleFunc("/api/tipos_transaccion/", api.CreateTipoTransaccion).Methods("POST")
}

func empleadosEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/empleados/", api.GetEmpleados).Methods("GET")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.GetStatus).Methods("GET")
	//mux.HandleFunc("/api/v1/status/", handlers.CreateCuenta).Methods("POST")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.UpdateStatus).Methods("PUT")
	//mux.HandleFunc("/api/v1/status/{id:[0-9]+}", handlers.DeleteStatus).Methods("DELETE")
}

func transaccionEndpoints(mux *mux.Router) {
	mux.HandleFunc("/api/transacciones/", api.GetTransacciones).Methods("GET")
	mux.HandleFunc("/api/transacciones/transferencias/", api.DoTransferencia).Methods("POST")
	mux.HandleFunc("/api/transacciones/depositos/", api.DoDeposito).Methods("POST")
}


func application(mux *mux.Router) {
	mux.HandleFunc("/", app.Index)
	mux.HandleFunc("/login/", app.Login)
	mux.HandleFunc("/cliente/", app.Cliente)
	mux.HandleFunc("/cajero/", app.Cajero)
	mux.HandleFunc("/admin/", app.Admin)
}
