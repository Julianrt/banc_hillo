package models

var loginSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS logins(
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    id_empleado INTEGER NOT NULL,
    token_string TEXT NOT NULL UNIQUE,
    fecha_registro TEXT NOT NULL,
    sesion_activa INTEGER NOT NULL
);`

type Login struct {
    ID              int     `json:"id"`
    IDEmpleado      int     `json:"id_empleado"`
    TokenString     string  `json:"token_string"`
    FechaRegistro   string  `json:"fecha_registro"`
    SesionActiva    int     `json:"sesion_activa"`
}

type TokenResponse struct {
    TokenString     string  `json:"token_string"`
    IDEmpleado      int     `json:"id_empleado"`
    Username        string  `json:"username"`
}

func nuevoLogin(idEmpleado int, tokenString string) *Login {
    login := &Login{
        IDEmpleado:     idEmpleado,
        TokenString:    tokenString,
        FechaRegistro:  ObtenerFechaHoraActualString(),
        SesionActiva:   1,
    }
    return login
}

func CrearLogin(idEmpleado int) (*Login, error) {
    tokenString, _ := RandomHex(20)
    token := nuevoLogin(idEmpleado, tokenString)
    err := token.Guardar()
    return token, err
}

func getLogin(sqlQuery string, condicion interface{}) (*Login, error) {
    login := &Login{}
    rows, err := Query(sqlQuery, condicion)
    for rows.Next() {
        rows.Scan(&login.ID, &login.IDEmpleado, &login.TokenString, &login.FechaRegistro, &login.SesionActiva)
    }
    return login, err
}

func GetLoginByToken(token string) (*Login, error) {
    sqlQuery := "SELECT id, id_empleado, token_string, fecha_registro, sesion_activa FROM logins WHERE token_string=?;"
    return getLogin(sqlQuery, token)
}

func (login *Login) Guardar() error {
    if login.ID == 0 {
        return login.registrar()
    }
    return login.actualizar()
}

func (login *Login) registrar() error {
    sqlQuery := "INSERT INTO logins(id_empleado, token_string, fecha_registro, sesion_activa) VALUES(?,?,?,?);"
    loginID, err := InsertData(sqlQuery, login.IDEmpleado, login.TokenString, login.FechaRegistro, login.SesionActiva)
    login.ID = int(loginID)
    return err
}

func (login *Login) actualizar() error {
    sqlQuery := "UPDATE logins SET id_empleado=?, token_string=?, fecha_registro=?, sesion_activa=? WHERE id=?;"
    _, err := Exec(sqlQuery, login.IDEmpleado, login.TokenString, login.FechaRegistro, login.SesionActiva, login.ID)
    return err
}

func (login *Login) Eliminar() error {
    sqlQuery := "DELETE FROM logins WHERE id=?;"
    _, err := Exec(sqlQuery, login.ID)
    return err
}