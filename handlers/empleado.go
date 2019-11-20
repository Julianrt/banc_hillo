package handlers

import (
	//"encoding/json"
	"net/http"
	//"strconv"

	"../models"
	//"github.com/gorilla/mux"
)


func GetEmpleados(w http.ResponseWriter, r *http.Request) {
	models.SendData(w, models.GetEmpleados())
}
