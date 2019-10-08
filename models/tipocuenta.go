package models

type TipoCuenta struct {
	ID 					int
	NombreTipoCuenta 	string
}

var tipoCuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tipos_cuenta(
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	tipo_cuenta TEXT NOT NULL
);`
