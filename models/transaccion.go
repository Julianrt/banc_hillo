package models

var transaccionSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS transacciones(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    fecha TEXT NOT NULL,
    monto REAL NOT NULL,
    estado INTEGER NOT NULL,
    numero_tarjeta_origen TEXT,
    numero_tarjeta_destino TEXT NOT NULL,
    id_tipo_transaccion INTEGER NOT NULL
);`

type Transaccion struct {
    ID                      int     `json:"id"`
    Fecha                   string  `json:"fecha"`
    Monto                   float32 `json:"monto"`
    Estado                  int     `json:"estado"`
    NumeroTarjetaOrigen     string  `json:"numero_tarjeta_origen"`
    NumeroTarjetaDestino    string  `json:"numero_tarjeta_destino"`
    IDTipoTransaccion       int     `json:"id_tipo_transaccion"`
}

type Transacciones []Transaccion

func NuevaTransaccion(monto float32, estado int, numeroTarjetaOrigen, numeroTarjetaDestino string, idTipoTransaccion int) *Transaccion {
    transaccion := &Transaccion{
        Fecha:                  ObtenerFechaHoraActualString(),
        Monto:                  monto,
        Estado:                 estado,
        NumeroTarjetaOrigen:    numeroTarjetaOrigen,
        NumeroTarjetaDestino:   numeroTarjetaDestino,
        IDTipoTransaccion:      idTipoTransaccion,
    }
    return transaccion
}

func CrearTransaccion(monto float32, estado int, numeroTarjetaOrigen, numeroTarjetaDestino string, idTipoTransaccion int) (*Transaccion, error) {
    transaccion := NuevaTransaccion(monto, estado, numeroTarjetaOrigen, numeroTarjetaDestino, idTipoTransaccion)
    err := transaccion.Guardar()
    return transaccion, err
}

func getTransaccion(query string, codicion interface{}) (*Transaccion, error) {
    transaccion := &Transaccion{}
    rows, err := Query(query, codicion)
    for rows.Next() {
        rows.Scan(&transaccion.ID, &transaccion.Fecha, &transaccion.Monto, &transaccion.Estado, 
            &transaccion.NumeroTarjetaOrigen, &transaccion.NumeroTarjetaDestino, &transaccion.IDTipoTransaccion)
    }
    return transaccion, err
}

func GetTransaccionByID(id int) (*Transaccion, error) {
    query := "SELECT id, fecha, monto, estado, numero_tarjeta_origen, numero_tarjeta_destino, id_tipo_transaccion FROM transacciones WHERE id=?"
    return getTransaccion(query, id)
}

func GetTransacciones() (Transacciones, error) {
    var transacciones Transacciones
    query := "SELECT id, fecha, monto, estado, numero_tarjeta_origen, numero_tarjeta_destino, id_tipo_transaccion FROM transacciones"
    rows, err := Query(query)
    for rows.Next() {
        transaccion := Transaccion{}
        rows.Scan(&transaccion.ID, &transaccion.Fecha, &transaccion.Monto, &transaccion.Estado, 
            &transaccion.NumeroTarjetaOrigen, &transaccion.NumeroTarjetaDestino, &transaccion.IDTipoTransaccion)
        transacciones = append(transacciones, transaccion)
    }
    return transacciones, err
}

func (transaccion *Transaccion) Guardar() error {
    if transaccion.ID == 0 {
        return transaccion.registrar()
    }
    return transaccion.actualizar()
}

func (transaccion *Transaccion) registrar() error {
    query := "INSERT INTO transacciones(fecha, monto, estado, numero_tarjeta_origen, numero_tarjeta_destino, id_tipo_transaccion) VALUES(?,?,?,?,?,?);"
    transaccionID, err := InsertData(query, transaccion.Fecha, transaccion.Monto, transaccion.Estado,
        transaccion.NumeroTarjetaOrigen, transaccion.NumeroTarjetaDestino, transaccion.IDTipoTransaccion )
    transaccion.ID = int(transaccionID)
    return err
}

func (transaccion *Transaccion) actualizar() error {
    query := "UPDATE transacciones SET fecha=?, monto=?, estado=?, numero_tarjeta_origen=?, numero_tarjeta_destino=?, id_tipo_transaccion=? WHERE id=?"
    _, err := Exec(query, transaccion.Fecha, transaccion.Monto, transaccion.Estado, transaccion.NumeroTarjetaOrigen, 
        transaccion.NumeroTarjetaDestino, transaccion.IDTipoTransaccion, transaccion.ID)
    return err
}

func (transaccion *Transaccion) Eliminar() error {
    query := "DELETE FROM transacciones WHERE id=?"
    _, err := Exec(query, transaccion.ID)
    return err
}
