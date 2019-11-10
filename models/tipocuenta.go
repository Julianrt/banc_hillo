package models

var tipoCuentaSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tipos_cuenta(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tipo_cuenta TEXT NOT NULL
);`

type TipoCuenta struct {
    ID                 int
    NombreTipoCuenta   string
}

type TiposCuenta []TipoCuenta

func NuevoTipoCuenta(nombreTipoCuente string) *TipoCuenta {
    tipoCuenta := &TipoCuenta{ 
        NombreTipoCuenta:   nombreTipoCuente,
    }
    return tipoCuenta
}

func CrearTipoCuenta(nombreTipoCuente string) (*TipoCuenta, error) {
    tipoCuenta := NuevoTipoCuenta(nombreTipoCuente)
    err := tipoCuenta.Guardar()
    return tipoCuenta, err
}

func getTipoCuenta(query string, condicion interface{}) (*TipoCuenta, error) {
    tipoCuenta := &TipoCuenta{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&tipoCuenta.ID, &tipoCuenta.NombreTipoCuenta)
    }
    return tipoCuenta, err
}

func GetTipoCuentaByID(id int) (*TipoCuenta, error) {
    query := "SELECT id, tipo_cuenta FROM tipos_cuenta WHERE id=?"
    return getTipoCuenta(query, id)
}

func GetTiposCuenta() (TiposCuenta, error) {
    var tiposCuenta TiposCuenta
    query := "SELECT id, tipo_cuenta FROM tipos_cuenta"
    rows, err := Query(query)
    for rows.Next() {
        var tipoCuenta TipoCuenta
        rows.Scan(&tipoCuenta.ID, &tipoCuenta.NombreTipoCuenta)
        tiposCuenta = append(tiposCuenta, tipoCuenta)
    }
    return tiposCuenta, err
}

func (tipoCuenta *TipoCuenta) Guardar() error {
    if tipoCuenta.ID == 0 {
        return tipoCuenta.registrar()
    }
    return tipoCuenta.actualizar()
}

func (tipoCuenta *TipoCuenta) registrar() error {
    query := "INSERT INTO tipos_cuenta(tipo_cuenta) VALUES(?);"
    tipoCuentaID, err :=  InsertData(query, tipoCuenta.NombreTipoCuenta)
    tipoCuenta.ID = int(tipoCuentaID)
    return err
}

func (tipoCuenta *TipoCuenta) actualizar() error {
    query := "UPDATE tipos_cuenta SET tipo_cuenta=? WHERE id=?"
    _, err := Exec(query, tipoCuenta.NombreTipoCuenta, tipoCuenta.ID)
    return err
}

func (tipoCuenta *TipoCuenta) Eliminar() error {
    query := "DELETE FROM tipos_cuenta WHERE id=?"
    _, err := Exec(query, tipoCuenta.ID)
    return err
}
