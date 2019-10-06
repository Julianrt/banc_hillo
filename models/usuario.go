package models

type Usuario struct {
	Nombre 		string 	`json:"nombre"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
}

func (usuario *Usuario) CrearCuenta(cuenta Cuenta) {
	
}