package handlers

import (
	"encoding/json"
	"net/http"
	//"strconv"

	"../models"
	//"github.com/gorilla/mux"
)

type Transferencia struct {
	TarjetaOrigen 	string 		`json:"tarjeta_origen"`
	TarjetaDestino 	string  	`json:"tarjeta_destino"`
	Monto 			float32 	`json:"monto"`
}

/*func GetEmpleados(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetEmpleados())
}*/

/*
func GetCiudad(w http.ResponseWriter, r *http.Request) {
	if ciudad, err := getCiudadByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if ciudad.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, ciudad)
	}
}*/

//CreateCiudad method
func DoTransferencia(w http.ResponseWriter, r *http.Request) {
	var transferencia Transferencia
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&transferencia); err != nil {
		models.SendUnprocessableEntity(w)
	}
	tarjetaOrigen, err := models.GetTarjetaByNumeroTarjeta(transferencia.TarjetaOrigen)
	cuentaOrigen, err := models.GetCuentaByID(tarjetaOrigen.IDCuenta)
	
	tarjetaDestino, err := models.GetTarjetaByNumeroTarjeta(transferencia.TarjetaDestino)
	cuentaDestino, err := models.GetCuentaByID(tarjetaDestino.IDCuenta)
	
	err = cuentaOrigen.Transferir(cuentaDestino.NumeroDeCuenta, transferencia.Monto)
	if err != nil {
		models.SendNotFound(w)
		return
	}
		transaccion,_ := models.CrearTransaccion(transferencia.Monto, 1, transferencia.TarjetaOrigen, transferencia.TarjetaDestino, 2)
		models.SendData(w, transaccion)

}

//UpdateCiudad method
/*func UpdateCiudad(w http.ResponseWriter, r *http.Request) {
	ciudad, err := getCiudadByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	var ciudadResponse models.Ciudad
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ciudadResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	ciudadResponse.ID = ciudad.ID
	ciudadResponse.Save()
	models.SendData(w, ciudadResponse)
}

//DeleteCiudad method
func DeleteCiudad(w http.ResponseWriter, r *http.Request) {
	if ciudad, err := getCiudadByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		ciudad.Delete()
		models.SendNoContent(w)
	}
}

func getCiudadByRequest(r *http.Request) (*models.Ciudad, error) {
	vars := mux.Vars(r)
	ciudadID, _ := strconv.Atoi(vars["id"])

	ciudad, err := models.GetCiudad(ciudadID)
	if err != nil {
		return ciudad, err
	}
	return ciudad, nil
}*/
