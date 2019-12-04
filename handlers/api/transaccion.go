package api

import (
	"encoding/json"
	"net/http"
	//"strconv"

	"../../models"
	//"github.com/gorilla/mux"
)

type Transferencia struct {
	TarjetaOrigen 	    string 		`json:"tarjeta_origen"`
    FechaVencimiento    string      `json:"fecha_vencimiento"`
    Cvv                 string      `json:"cvv"`
	TarjetaDestino      string  	`json:"tarjeta_destino"`
	Monto               float32 	`json:"monto"`
}

type Desposito struct {
	TarjetaDestino 	string  	`json:"tarjeta_destino"`
	Monto 			float32 	`json:"monto"`
}

func GetTransacciones(w http.ResponseWriter, r *http.Request) {
	transacciones,_ := models.GetTransacciones()
	models.SendData(w, transacciones)
}

func DoTransferencia(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var transferencia Transferencia
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&transferencia); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}

    if models.ValidTarjeta(transferencia.TarjetaOrigen, transferencia.FechaVencimiento, transferencia.Cvv) {

        tarjetaOrigen, err := models.GetTarjetaByNumeroTarjeta(transferencia.TarjetaOrigen)
        cuentaOrigen, err := models.GetCuentaByID(tarjetaOrigen.IDCuenta)
    	
        tarjetaDestino, err := models.GetTarjetaByNumeroTarjeta(transferencia.TarjetaDestino)
    	cuentaDestino, err := models.GetCuentaByID(tarjetaDestino.IDCuenta)

    	if tarjetaOrigen.ID != 0 && tarjetaDestino.ID != 0 {
        	err = cuentaOrigen.Transferir(cuentaDestino.NumeroDeCuenta, transferencia.Monto)
        	if err != nil {
            	models.SendNotFound(w)
            	return
        	}
        	transaccion,_ := models.CrearTransaccion(transferencia.Monto, 1, transferencia.TarjetaOrigen, transferencia.TarjetaDestino, 2)
        	models.SendData(w, transaccion)
        	return
    	}
    }

    models.SendNotFound(w)
}

func DoDeposito(w http.ResponseWriter, r *http.Request) {
	var deposito Desposito
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&deposito); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	tarjetaDestino, err := models.GetTarjetaByNumeroTarjeta(deposito.TarjetaDestino)
	cuentaDestino, err := models.GetCuentaByID(tarjetaDestino.IDCuenta)
	
	err = cuentaDestino.Depositar(deposito.Monto)
	if err != nil {
		models.SendNotFound(w)
		return
	}
	
	transaccion,_ := models.CrearTransaccion(deposito.Monto, 1, "", deposito.TarjetaDestino, 1)
	models.SendData(w, transaccion)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
