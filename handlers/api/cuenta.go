package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
)


func GetCuentas(w http.ResponseWriter, r *http.Request) {
	cuentas,_ := models.GetCuentas()
	models.SendData(w, cuentas)
}

func GetCuenta(w http.ResponseWriter, r *http.Request) {
	if cuenta, err := getCuentaByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if cuenta.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, cuenta)
	}
}

func CreateCuenta(w http.ResponseWriter, r *http.Request) {
	var cuenta models.Cuenta
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cuenta); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		cuenta.Guardar()
		models.SendData(w, cuenta)
	}
}

func getCuentaByRequest(r *http.Request) (*models.Cuenta, error) {
	vars := mux.Vars(r)
	cuentaID, _ := strconv.Atoi(vars["id"])

	cuenta, err := models.GetCuentaByID(cuentaID)
	return cuenta, err
}
