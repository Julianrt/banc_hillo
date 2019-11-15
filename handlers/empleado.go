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
}

//CreateCiudad method
func CreateCiudad(w http.ResponseWriter, r *http.Request) {
	var ciudad models.Ciudad
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&ciudad); err != nil {
		models.SendUnprocessableEntity(w)
	} else {
		ciudad.Save()
		models.SendData(w, ciudad)
	}
}

//UpdateCiudad method
func UpdateCiudad(w http.ResponseWriter, r *http.Request) {
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
