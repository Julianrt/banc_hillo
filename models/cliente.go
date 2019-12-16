package models

var clienteSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS clientes(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    nombre TEXT NOT NULL,
    apellido_paterno TEXT,
    apellido_materno TEXT,
    clave TEXT UNIQUE,
    habilitado INTEGER NOT NULL,
    fecha_creacion TEXT NOT NULL
);`

type Cliente struct {
    ID              int     `json:"id"`
    Nombre          string  `json:"nombre"`
    ApellidoPaterno string  `json:"apellido_paterno"`
    ApellidoMaterno string  `json:"apellido_materno"`
    Clave           string  `json:"clave"`
    habilitado      int
    fechaCreacion   string
}

type Clientes []Cliente

func NuevoCliente(nombre, apellido_paterno, apellido_materno, clave string) *Cliente {
    cliente := &Cliente{
        Nombre:             nombre,
        ApellidoPaterno:    apellido_paterno,
        ApellidoMaterno:    apellido_materno,
        Clave:              clave,
        habilitado:         1,
        fechaCreacion:      ObtenerFechaHoraActualString(),
    }
    return cliente
}

func CrearCliente(nombre, apellido_paterno, apellido_materno, clave string) (*Cliente, error) {
    cliente := NuevoCliente(nombre, apellido_paterno, apellido_materno, clave)
    err := cliente.Guardar()
    return cliente, err
}

func GetClientes() (Clientes, error) {
    var clientes Clientes
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1"
    rows, err := Query(query)
    for rows.Next(){
        cliente := Cliente{}
        rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.ApellidoPaterno, &cliente.ApellidoMaterno, &cliente.Clave,
            &cliente.habilitado, &cliente.fechaCreacion)
        clientes = append(clientes, cliente)
    }
    return clientes, err
}

func getCliente(query string, condicion interface{}) (*Cliente, error) {
    cliente := Cliente{}
    rows, err := Query(query, condicion)
    for rows.Next() {
        rows.Scan(&cliente.ID, &cliente.Nombre, &cliente.ApellidoPaterno, &cliente.ApellidoMaterno, &cliente.Clave,
            &cliente.habilitado, &cliente.fechaCreacion)
    }
    return &cliente, err
}

func GetClienteByID(id int) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND id=?"
    return getCliente(query, id)
}

func GetClienteByClave(clave string) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND clave=?"
    return getCliente(query, clave)
}

func GetClienteByNumeroTarjeta(tarjeta string) (*Cliente, error) {
    query := "SELECT id, nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion FROM clientes WHERE habilitado=1 AND id in (SELECT id_cliente FROM tarjetas WHERE numero_tarjeta=?)"
    return getCliente(query, tarjeta)
}

func (cliente *Cliente) Guardar() error {
    if cliente.ID == 0 {
        return cliente.registrar()
    }
    return cliente.actualizar()
}

func (cliente *Cliente) registrar() error {
    cliente.habilitado=1
    cliente.fechaCreacion=ObtenerFechaHoraActualString()
    query := "INSERT INTO clientes(nombre, apellido_paterno, apellido_materno, clave, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?);"
    clienteID, err := InsertData(query, cliente.Nombre, cliente.ApellidoPaterno, cliente.ApellidoMaterno, cliente.Clave,
        cliente.habilitado, cliente.fechaCreacion)
    cliente.ID = int(clienteID)
    return err
}

func (cliente *Cliente) actualizar() error {
    query := "UPDATE clientes SET nombre=?, apellido_paterno=?, apellido_materno=?, clave=?, habilitado=? WHERE id=?"
    _, err := Exec(query, cliente.Nombre, cliente.ApellidoPaterno, cliente.ApellidoMaterno, cliente.habilitado, cliente.ID )
    return err
}

func (cliente *Cliente) EliminarLog() error {
    cliente.habilitado=0
    return cliente.actualizar()
}

func (cliente *Cliente) Eliminar() error {
    query := "DELETE FROM clientes WHERE id=?"
    _, err := Exec(query, cliente.ID)
    return err
}
