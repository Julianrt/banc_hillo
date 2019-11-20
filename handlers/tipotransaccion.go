package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../models"
	"github.com/gorilla/mux"
)


func GetTiposTransaccion(w http.ResponseWriter, r *http.Request) {
	tiposTransaccion,_ := models.GetTiposTransaccion()
	models.SendData(w, tiposTransaccion)
}

func GetTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	if tipoTransaccion, err := getTiposTransaccionByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if tipoTransaccion.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, tipoTransaccion)
	}
}

func CreateTipoTransaccion(w http.ResponseWriter, r *http.Request) {
	var tipoTransaccion models.TipoTransaccion
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tipoTransaccion); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		tipoTransaccion.Guardar()
		models.SendData(w, tipoTransaccion)
	}
}

func getTiposTransaccionByRequest(r *http.Request) (*models.TipoTransaccion, error) {
	vars := mux.Vars(r)
	tipoTransaccionID, _ := strconv.Atoi(vars["id"])

	tipoTransaccion, err := models.GetTipoTransaccionByID(tipoTransaccionID)
	return tipoTransaccion, err
}
