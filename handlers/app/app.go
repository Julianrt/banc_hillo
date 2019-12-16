package app

import (
	"net/http"
	"../../utils"
	"log"
	"strconv"
	"../../models"
)

type ServeToClient struct {
	Transacciones 	models.Transacciones
	Cuenta 			*models.Cuenta
	Tarjetas 		models.Tarjetas
}

func Index(w http.ResponseWriter, r *http.Request) {


	//conext := make(map[string]interface{})
	//conext["Authenticated"] = utils.IsAuthenticated(r)

	utils.RenderTemplate(w, "app/index", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if _,err := strconv.Atoi(username); err == nil {
			tarjeta, _ := models.GetTarjetaByNumeroTarjeta(username)
			if tarjeta.ID != 0 {

				if tarjeta.Nip == password {

					http.Redirect(w, r, "/cliente/?token="+username, 302)
					return

				} else {
					log.Println("password no coincide")
				}

			} else {
				log.Println("No se encontro la tarjeta")
			}
		} else {
			log.Println("no es una tarjeta")
			empleado,_ := models.GetEmpleadoByUsername(username)
			if empleado.ID != 0{
				if empleado.Password == password {
					http.Redirect(w, r, "/cajero/", 302)
					return
				} else {
					log.Println("password no coincide")
				}
			} else {
				log.Println("usuario incorrecto del empleado")
			}
		}
		http.Redirect(w, r, "/login/", 302)

	} else if r.Method == "GET" {
		utils.RenderTemplate(w, "app/login", nil)
	}

}

func Cliente(w http.ResponseWriter, r *http.Request) {


	if r.Method == "GET" {
		token := r.URL.Query().Get("token")
		if token != "" {
			cuenta, _ := models.GetCuentaByNumeroTarjeta(token)
			transacciones,_ := models.GetTransaccionesByTerjeta(token)
			tarjetas,_ := models.GetTarjetasByIDCuenta(cuenta.ID)

			response := ServeToClient{transacciones, cuenta, tarjetas}

			utils.RenderTemplate(w, "app/cliente", response)
		}
	} else if r.Method == "POST" {
		numeroTarjeta := r.FormValue("tarjeta")
		http.Redirect(w, r, "/cliente/?token="+numeroTarjeta, 302)
	}
}

func Cajero(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.RenderTemplate(w, "app/cajero", nil)
	} else if r.Method == "POST" {
		numeroTarjetaOrigen := r.FormValue("tarjetaOrigen")
		cvv := r.FormValue("cvv");
		mes:= r.FormValue("mes");
		ano:= r.FormValue("ano");
		numeroTarjetaDestino:= r.FormValue("tarjetaDestino");
		montoString:= r.FormValue("monto");
		monto := convMonto(montoString)
		fechaVencimiento := mes+"/"+ano

		cuentaDestino,_ := models.GetCuentaByNumeroTarjeta(numeroTarjetaDestino)

		if numeroTarjetaOrigen != "" {
			//Transferencia
			tarjetaOrigen, _ := models.GetTarjetaByNumeroTarjeta(numeroTarjetaOrigen)
			log.Println("Transferencia")
			if tarjetaOrigen.ID != 0 {
				if tarjetaOrigen.FechaVencimiento == fechaVencimiento && tarjetaOrigen.NumeroSeguridad == cvv {
					cuentaOrigen,_ := models.GetCuentaByID(tarjetaOrigen.IDCuenta)
					if cuentaDestino.ID != 0 {
						cuentaOrigen.Transferir(cuentaDestino.NumeroDeCuenta, monto)
						models.CrearTransaccion(monto,1,numeroTarjetaOrigen,numeroTarjetaDestino,2)
					}else {
						log.Println("Ese tarjeta NO tiene cuenta")
					}
				} else {
					log.Println("Los datos de la tarjeta de origen no coinciden")
				}
			} else {
				log.Println("No existe esa tarjeta de origen")
			}
		} else {
			//Deposito
			log.Println("Deposito")
			if cuentaDestino.ID != 0 {
				cuentaDestino.Depositar(monto)
				models.CrearTransaccion(monto,1,"",numeroTarjetaDestino,1)
			} else {
				log.Println("Esta tarjeta no tiene cuenta")
			}
		}


		http.Redirect(w, r, "/cajero/", 302)

	}

}

func Admin(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		utils.RenderTemplate(w, "app/admin", nil)
	} else if r.Method == "POST" {
		
	}
}

func convMonto(montoString string) float32 {
	monto,_ := strconv.ParseFloat(montoString, 32);
	return float32(monto)
}
