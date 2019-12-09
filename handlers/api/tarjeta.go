package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
)


func GetTarjetas(w http.ResponseWriter, r *http.Request) {
	tarjetas,_ := models.GetTarjetas()
	models.SendData(w, tarjetas)
}

func GetTarjeta(w http.ResponseWriter, r *http.Request) {
	if tarjeta, err := getTarjetaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if tarjeta.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, tarjeta)
	}
}

func CreateTarjeta(w http.ResponseWriter, r *http.Request) {
	var tarjeta models.Tarjeta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tarjeta); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		tarjeta.Guardar()
		models.SendData(w, tarjeta)
	}
}

func UpdateTarjeta(w http.ResponseWriter, r *http.Request) {
	tarjeta, err := getTarjetaByRequest(r)
	if err != nil {
		models.SendNotFound(w)
		return
	}

	var tarjetaResponse models.Tarjeta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&tarjetaResponse); err != nil {
		models.SendUnprocessableEntity(w)
		return
	}
	tarjetaResponse.ID = tarjeta.ID
	tarjetaResponse.Guardar()
	models.SendData(w, tarjetaResponse)
}

func getTarjetaByRequest(r *http.Request) (*models.Tarjeta, error) {
	vars := mux.Vars(r)
	tarjetaID, _ := strconv.Atoi(vars["id"])

	tarjeta, err := models.GetTarjetaByID(tarjetaID)
	return tarjeta, err
}
