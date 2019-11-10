package models

var tipoTransaccionSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS tipos_transaccion(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    tipo_transaccion TEXT NOT NULL
);`

type TipoTransaccion struct {
    ID                      int     `json:"id"`
    NombreTipoTransaccion   string  `json:"tipo_transaccion"`
}

type TiposTransaccion []TipoTransaccion

func NuevoTipoTransaccion (nombreTipoTransaccion string) *TipoTransaccion {
    tipoTransaccion := &TipoTransaccion{
        NombreTipoTransaccion:  nombreTipoTransaccion,
    }
    return tipoTransaccion
}

func CrearTipoTransaccion (nombreTipoTransaccion string) (*TipoTransaccion, error) {
    tipoTransaccion := NuevoTipoTransaccion(nombreTipoTransaccion)
    err := tipoTransaccion.Guardar()
    return tipoTransaccion, err
}

func GetTipoTransaccion(query string, condicion interface{}) (*TipoTransaccion, error) {
    tipoTransaccion := &TipoTransaccion{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&tipoTransaccion.ID, &tipoTransaccion.NombreTipoTransaccion)
    }
    return tipoTransaccion, err
}

func GetTipoTransaccionByID (id int) (*TipoTransaccion, error) {
    query := "SELECT id, tipo_transaccion FROM tipos_transaccion WHERE id = ?"
    return GetTipoTransaccion(query, id)
}

func GetTiposTransaccion() (TiposTransaccion, error) {
    var tiposTransaccion TiposTransaccion
    query := "SELECT id, tipo_transaccion FROM tipos_transaccion"
    rows, err := Query(query)
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        tipoTransaccion := TipoTransaccion{}
        rows.Scan(&tipoTransaccion.ID, &tipoTransaccion.NombreTipoTransaccion )
        tiposTransaccion = append(tiposTransaccion, tipoTransaccion)
    }
    return tiposTransaccion, err
}

func (tipoTransaccion *TipoTransaccion) Guardar() error {
    if tipoTransaccion.ID == 0 {
        return tipoTransaccion.registrar()
    }
    return tipoTransaccion.actualizar()
}

func (tipoTransaccion *TipoTransaccion) registrar() error {
    query := "INSERT INTO tipos_transaccion (tipo_transaccion) VALUES(?)"
    tipoTransaccionID, err := InsertData(query, tipoTransaccion.NombreTipoTransaccion)
    tipoTransaccion.ID = int(tipoTransaccionID)
    return err
}

func (tipoTransaccion *TipoTransaccion) actualizar() error {
    query := "UPDATE tipos_transaccion SET tipo_transaccion=? WHERE id=?"
    _, err := Exec(query, tipoTransaccion.NombreTipoTransaccion, tipoTransaccion.ID)
    return err
} 

func (tipoTransaccion *TipoTransaccion) Eliminar() error {
    query := "DELETE FROM tipos_transaccion WHERE id=?"
    _, err := Exec(query, tipoTransaccion.ID)
    return err
}
