package models

type Empleado struct {
    ID                 int       `json:"id"`
    Nombre             string    `json:"nombre"`
    ApellidoPaterno    string    `json:"apellido_paterno"`
    ApellidoMaterno    string    `json:"apellido_materno"` 
    Username           string    `json:"username"`
    Password           string    `json:"password"`
    habilitado         int
    fechaCreacion      string
}

var empleadosSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS empleados(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nombre TEXT NOT NULL,
	apellido_paterno TEXT NOT NULL,
	apellido_materno TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	habilitado INTEGER NOT NULL,
	fecha_creacion TEXT NOT NULL
);`

type Empleados []Empleado

func NuevoEmpleado(nombre, apellido_paterno, apellido_materno, username, password string) *Empleado {
    empleado := &Empleado{
        Nombre:	          nombre,
        ApellidoPaterno:  apellido_paterno,
        ApellidoMaterno:  apellido_materno,
        Username:         username,
        Password:         password,
        habilitado:       1,
        fechaCreacion:    ObtenerFechaHoraActualString(),
    }
    return empleado
}

func CrearEmpleado(nombre, apellido_paterno, apellido_materno, username, password string) (*Empleado, error) {
	empleado := NuevoEmpleado(nombre, apellido_paterno, apellido_materno, username, password)
	err := empleado.Save()
	return empleado, err
}

func GetEmpleado(sql string, condicion interface{}) (*Empleado, error) {
    empleado := &Empleado{}
    rows, err := Query(sql, condicion)
    for rows.Next() {
        rows.Scan(&empleado.ID, &empleado.Nombre, &empleado.ApellidoPaterno, &empleado.ApellidoMaterno, 
            &empleado.Username, &empleado.Password, &empleado.habilitado, &empleado.fechaCreacion)
    }
    return empleado, err
}

func GetEmpleadoByID(id int) (*Empleado, error) {
    sql := "SELECT id, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1 AND id=?"
    return GetEmpleado(sql, id)
}

func GetEmpleadoByUsername(username string) (*Empleado, error) {
    sql := "SELECT id, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1 AND username=?"
    return GetEmpleado(sql, username)
}

func GetEmpleados() Empleados {
    var empleados Empleados
    sql := "SELECT id, nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion FROM empleados WHERE habilitado=1"
    rows, _ := Query(sql)
    for rows.Next() {
        var empleado Empleado
        rows.Scan(&empleado.ID, &empleado.Nombre, &empleado.ApellidoPaterno, &empleado.ApellidoMaterno, 
            &empleado.Username, &empleado.Password, &empleado.habilitado, &empleado.fechaCreacion)
        empleados = append(empleados, empleado)
    }
    return empleados
}

/*func LoginEmpleado(username, password string) (*Empleado, error) {
	empleado,_ := GetEmpleadoByUsername(username)
	if empleado.Password != password {
		return &Empleado{}, errors.New("Usuario o contraseña no coinciden")
	}
	return empleado, nil
}*/

func (empleado *Empleado) Save() error {
	if empleado.ID == 0 {
		return empleado.registrar()
	}
	return empleado.actualizar()
}

func (empleado *Empleado) registrar() error {
	sql := "INSERT INTO empleados(nombre, apellido_paterno, apellido_materno, username, password, habilitado, fecha_creacion) VALUES(?,?,?,?,?,?,?);"
	empleadoID, err := InsertData(sql, empleado.Nombre, empleado.ApellidoPaterno, empleado.ApellidoMaterno, 
        empleado.Username, empleado.Password, empleado.habilitado, empleado.fechaCreacion)
	empleado.ID = int(empleadoID)
	return err
}

func (empleado *Empleado) actualizar() error {
	sql := "UPDATE empleados SET nombre=?, apellido_paterno=?, apellido_materno=?, username=?, password=?, habilitado=? WHERE id=?"
	_, err := Exec(sql, empleado.Nombre, empleado.ApellidoPaterno, empleado.ApellidoMaterno, empleado.Username, empleado.Password, 
        empleado.habilitado, empleado.ID)
	return err
}

func (empleado *Empleado) EliminarLog() error {
    empleado.habilitado=0
    return empleado.actualizar()
}

func (empleado *Empleado) Eliminar() error {
	sql := "DELETE FROM empleados WHERE id=?"
	_, err := Exec(sql, empleado.ID)
	return err
}

func (empleado *Empleado) SetPassword(password string) {
	empleado.Password = password
    empleado.Save()
}
