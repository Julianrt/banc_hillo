package models

type Cuenta struct {
	ID 				int 		`json:"id"`
	NumeroDeCuenta	int 		`json:"numero_de_cuenta"`
	Nip				int 		`json:"nip"`
	Saldo			double 		`json:"saldo"`
	Titular			string 		`json:"titular"`
	IDTipoDeCuenta 	int 		`json:"id_tipo_de_cuenta"`
}

cuentaSchemeSQLITE := `CREATE TABLE IF NOT EXISTS cuentas(
	id INTEGER PRIMARY KEY,
	numero_de_cuenta INTEGER NOT NULL,
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

func (cuenta *Cuenta) Depositar(monto double) {

}


func (cuenta *Cuenta) Retirar(monto double) {

}

func (cuenta *Cuenta) Transferir(cuentaDestino int, monto double) {

}

func (cuenta *Cuenta) SolicitarSaldo() (double) {

}

func (cuenta *Cuenta) Guardar() error {
	if cuenta.ID == 0 {
		return cuenta.registrar()
	} 

	return cuenta.actualizar()
}

func (cuenta *Cuenta) registrar() error {
	query := "INSERTO INTO cuentas(numero_de_cuenta, nip, saldo, titular, id_tipo_de_cuenta) VALUES(?,?,?,?,?);"
}
