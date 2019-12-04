package app

import (
	"net/http"
	"../../utils"
)

func Index(w http.ResponseWriter, r *http.Request) {


	//conext := make(map[string]interface{})
	//conext["Authenticated"] = utils.IsAuthenticated(r)

	utils.RenderTemplate(w, "app/index", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		http.Redirect(w, r, "/cliente/", 302)

	} else if r.Method == "GET" {
		utils.RenderTemplate(w, "app/login", nil)
	}

}

func Cliente(w http.ResponseWriter, r *http.Request) {


	if r.Method == "GET" {
		utils.RenderTemplate(w, "app/cliente", nil)
	} else if r.Method == "POST" {

	}

}