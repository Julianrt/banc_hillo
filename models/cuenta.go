package models

import "errors"

type Cuenta struct {
	ID 				int 		`json:"id"`
	NumeroDeCuenta 	int 		`json:"numero_de_cuenta"`
	Nip 			int 		`json:"nip"`
	Saldo 			float32 	`json:"saldo"`
	Titular 		string 		`json:"titular"`
	IDTipoDeCuenta 	int 		`json:"id_tipo_de_cuenta"`
}

var cuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS cuentas(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	numero_de_cuenta INTEGER NOT NULL UNIQUE,
	nip INTEGER NOT NULL,
	saldo REAL,
	titular TEXT,
	id_tipo_de_cuenta INTEGER);`


func nuevaCuenta(numeroDeCuenta, nip int, titular string, idTipoDeCuenta int) *Cuenta {
	cuenta := &Cuenta {
		NumeroDeCuenta: numeroDeCuenta,
		Nip: 			nip,
		Saldo:			0.0,
		Titular:		titular,
		IDTipoDeCuenta:	idTipoDeCuenta,
	}
	return cuenta
}

func AltaCuenta(numeroDeCuenta, nip int, titular string, idTipoDeCuenta int) (*Cuenta, error) {
	cuenta := nuevaCuenta(numeroDeCuenta, nip, titular, idTipoDeCuenta )
	err := cuenta.Guardar()
	return cuenta, err
}

func GetCuentaByNumeroCuenta(numeroDeCuenta int) (*Cuenta, error) {
	cuenta := nuevaCuenta(0,0,"",0)
	query := "SELECT id, numero_de_cuenta, nip, saldo, titular, id_tipo_de_cuenta FROM cuentas WHERE numero_de_cuenta = ?"
	rows, err := Query(query, numeroDeCuenta)
	if err != nil {
		return cuenta, err
	}
	for rows.Next() {
		rows.Scan(&cuenta.ID, &cuenta.NumeroDeCuenta, &cuenta.Nip, &cuenta.Saldo, &cuenta.Titular, &cuenta.IDTipoDeCuenta)
	}
	return cuenta, nil
}

func (cuenta *Cuenta) Depositar(monto float32) error {
	cuenta.Saldo = monto
	return cuenta.Guardar()
}

func (cuenta *Cuenta) Retirar(monto float32) error {
	if cuenta.Saldo >= monto {
		cuenta.Saldo -= monto
		return cuenta.Guardar()
	} else {
		return errors.New("Saldo insuficiente")
	}
}

func (cuenta *Cuenta) Transferir(numeroCuentaDestino int, monto float32) error {
	cuentaDestino := nuevaCuenta(0,0,"",0)
	err := errors.New("")

	if cuentaDestino, err = GetCuentaByNumeroCuenta(numeroCuentaDestino); err != nil {
		return err
	}
	if err = cuenta.Retirar(monto); err != nil {
		return err
	}
	err = cuentaDestino.Depositar(monto)
	return err
}

func (cuenta *Cuenta) SolicitarSaldo() (float32, error) {
	var saldo float32
	query := "SELECT saldo FROM cuentas WHERE numero_de_cuenta = ?"
	rows, err := Query(query, cuenta.NumeroDeCuenta)
	if err != nil {
		return saldo, err
	}
	for rows.Next() {
		rows.Scan(&saldo)
	}
	return saldo, nil
}

func (cuenta *Cuenta) Guardar() error {
	if cuenta.ID == 0 {
		return cuenta.registrar()
	} 

	return cuenta.actualizar()
}

func (cuenta *Cuenta) registrar() error {
	query := "INSERT INTO cuentas(numero_de_cuenta, nip, saldo, titular, id_tipo_de_cuenta) VALUES(?,?,?,?,?);"
	cuentaID, err := InsertData(query, cuenta.NumeroDeCuenta, cuenta.Nip, cuenta.Saldo, cuenta.Titular, cuenta.IDTipoDeCuenta)
	cuenta.ID = int(cuentaID)
	return err
}

func (cuenta *Cuenta) actualizar() error {
	query := "UPDATE cuentas SET nip=?, saldo=?, titular=?, id_tipo_de_cuenta=? WHERE numero_de_cuenta=?;"
	_, err := Exec(query, cuenta.Nip, cuenta.Saldo, cuenta.Titular, cuenta.IDTipoDeCuenta, cuenta.NumeroDeCuenta)
	return err
}
