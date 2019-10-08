package models

type Usuario struct {
	ID 			int 	`json:"id"`
	Nombre 		string 	`json:"nombre"`
	Username 	string 	`json:"username"`
	Password 	string 	`json:"password"`
}

var usuarioSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS usuarios(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nombre TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);`

func (usuario *Usuario) CrearCuenta(cuenta Cuenta) {
	
}
