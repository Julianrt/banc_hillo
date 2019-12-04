package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../../models"
	"github.com/gorilla/mux"
)


func GetClientes(w http.ResponseWriter, r *http.Request) {
	clientes,_ := models.GetClientes()
	models.SendData(w, clientes)
}

func GetCliente(w http.ResponseWriter, r *http.Request) {
	if cliente, err := getClienteByRequest(r); err != nil {
		models.SendNotFound(w)
	} else {
		if cliente.ID == 0 {
			models.SendNotFound(w)
			return
		}
		models.SendData(w, cliente)
	}
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var cliente models.Cliente
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&cliente); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		cliente.Guardar()
		models.SendData(w, cliente)
	}
}

func getClienteByRequest(r *http.Request) (*models.Cliente, error) {
	vars := mux.Vars(r)
	clienteID, _ := strconv.Atoi(vars["id"])

	cliente, err := models.GetClienteByID(clienteID)
	return cliente, err
}
