package app

import (
	"net/http"
	"../../utils"
	//"strconv"
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

func Cajero(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.RenderTemplate(w, "app/cajero", nil)
	} else if r.Method == "POST" {
		tarjetaOrigen := r.FormValue("tarjetaOrigen")
		cvv := r.FormValue("cvv");
		mes:= r.FormValue("mes");
		ano:= r.FormValue("ano");
		tarjetaDestino:= r.FormValue("tarjetaDestino");
		monto:= r.FormValue("monto");
		if tarjetaDestino != "" && monto != "" {
			//montito,_ := strconv.Atoi(monto)
			if tarjetaOrigen != "" && cvv != "" && mes != "" && ano!=""{
				//Aqui Transferencia
				
			}else{
				//Aqui Deposito

			}
		} 
	}

}

func Admin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.RenderTemplate(w, "app/admin", nil)
	} else if r.Method == "POST" {

	}
}