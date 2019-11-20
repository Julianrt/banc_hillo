package models

import "errors"

type Cuenta struct {
    ID                 int     `json:"id"`
    NumeroDeCuenta     string  `json:"numero_de_cuenta"`
    Saldo              float32 `json:"saldo"`
    IDCliente          int     `json:"id_cliente"`
    IDTipoDeCuenta     int     `json:"id_tipo_de_cuenta"`
    habilitado         int
    fechaCreacion      string
}

type Cuentas []Cuenta

var cuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS cuentas(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    numero_de_cuenta TEXT NOT NULL UNIQUE,
    saldo REAL DEFAULT 0.0,
    id_cliente TEXT,
    id_tipo_de_cuenta INTEGER,
    habilitado INTEGER DEFAULT 0,
    fecha_creacion TEXT);`


func nuevaCuenta(numeroDeCuenta string, idCliente, idTipoDeCuenta int) *Cuenta {
    cuenta := &Cuenta {
        NumeroDeCuenta: numeroDeCuenta,
        Saldo:          0.0,
        IDCliente:      idCliente,
        IDTipoDeCuenta: idTipoDeCuenta,
        habilitado:     0,
        fechaCreacion:  ObtenerFechaHoraActualString(),
    }
    return cuenta
}

func AltaCuenta(numeroDeCuenta string, idCliente, idTipoDeCuenta int) (*Cuenta, error) {
    cuenta := nuevaCuenta(numeroDeCuenta, idCliente, idTipoDeCuenta )
    err := cuenta.Guardar()
    return cuenta, err
}

func getCuenta(sqlQuery string, condicion interface{}) (*Cuenta, error) {
	cuenta := &Cuenta{}
	rows, err := Query(sqlQuery, condicion)
	for rows.Next() {
		rows.Scan(&cuenta.ID, &cuenta.NumeroDeCuenta, &cuenta.Saldo, &cuenta.IDCliente, &cuenta.IDTipoDeCuenta, 
            &cuenta.habilitado, &cuenta.fechaCreacion)
	}
	return cuenta, err
}

func GetCuentaByID(id int) (*Cuenta, error) {
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE id=?"
	return getCuenta(query, id)
}

func GetCuentaByNumeroCuenta(numeroDeCuenta string) (*Cuenta, error) {
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE numero_de_cuenta=?"
	return getCuenta(query, numeroDeCuenta)
}

func GetCuentas() (Cuentas, error) {
	var cuentas Cuentas
	query := "SELECT id, numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion FROM cuentas WHERE habilitado=1"
	rows, err := Query(query)
	for rows.Next() {
		cuenta := Cuenta{}
		rows.Scan(&cuenta.ID, &cuenta.NumeroDeCuenta, &cuenta.Saldo, &cuenta.IDCliente, &cuenta.IDTipoDeCuenta, 
            &cuenta.habilitado, &cuenta.fechaCreacion)
        cuentas = append(cuentas, cuenta)
	}
	return cuentas, err
}

func (cuenta *Cuenta) Depositar(monto float32) error {
	cuenta.Saldo += monto
	err := cuenta.Guardar()
	if err == nil {
		cuenta.activarCuenta()
	}
	return err
}

func (cuenta *Cuenta) Retirar(monto float32) error {
	if cuenta.Saldo >= monto {
		cuenta.Saldo -= monto
		return cuenta.Guardar()
	} else {
		return errors.New("Saldo insuficiente")
	}
}

func (cuenta *Cuenta) Transferir(numeroCuentaDestino string, monto float32) error {
	cuentaDestino := &Cuenta{}
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
	cuenta.fechaCreacion=ObtenerFechaHoraActualString()
	query := "INSERT INTO cuentas(numero_de_cuenta, saldo, id_cliente, id_tipo_de_cuenta, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?);"
	cuentaID, err := InsertData(query, cuenta.NumeroDeCuenta, cuenta.Saldo, cuenta.IDCliente, cuenta.IDTipoDeCuenta, 
        cuenta.habilitado, cuenta.fechaCreacion)
	cuenta.ID = int(cuentaID)
	return err
}

func (cuenta *Cuenta) actualizar() error {
	query := "UPDATE cuentas SET saldo=?, id_cliente=?, id_tipo_de_cuenta=?, habilitado=?, fecha_creacion=? WHERE numero_de_cuenta=?;"
	_, err := Exec(query, cuenta.Saldo, cuenta.IDCliente, cuenta.IDTipoDeCuenta, cuenta.habilitado, cuenta.fechaCreacion, 
        cuenta.NumeroDeCuenta)
	return err
}

func (cuenta *Cuenta) estaActivada() bool {	
	if cuenta.habilitado == 1 {
		return true
	}
	return false
}

func (cuenta *Cuenta) activarCuenta() error {
	var err error

	if !cuenta.estaActivada() {
		query := "UPDATE cuentas SET habilitado=1 WHERE id=?"
		_, err = Exec(query, cuenta.ID)
		if err == nil {
			cuenta.habilitado = 1
		}
	}
	return err
}

