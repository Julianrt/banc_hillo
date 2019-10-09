package models

import "errors"

type Usuario struct {
	ID         int     `json:"id"`
	Nombre     string  `json:"nombre"`
	Username   string  `json:"username"`
	Password   string  `json:"password"`
}

var usuarioSchemeSQLITE string = `CREATE TABLE IF NOT EXISTS usuarios(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	nombre TEXT NOT NULL,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
);`

type Usuarios []Usuario

func NuevoUsuario(nombre, username, password string) (*Usuario) {
	usuario := &Usuario{
		Nombre:		nombre,
		Username:	username,
		Password:	password,
	}
	return usuario
}

func CrearUsuario(nombre, username, password string) (*Usuario, error) {
	usuario := NuevoUsuario(nombre, username, password)
	err := usuario.Save()
	return usuario, err
}

func getUsuario(sql string, condicion interface{}) (*Usuario, error) {
	usuario := &Usuario{}
	rows, err := Query(sql, condicion)
	for rows.Next() {
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Password)
	}
	return usuario, err
}

func GetUsuarioByID(id int) (*Usuario, error) {
	sql := "SELECT id, nombre, username, password FROM usuarios WHERE id=?"
	usuario, err := getUsuario(sql, id)
	return usuario, err
}

func GetUsuarioByUsername(username string) (*Usuario, error) {
	sql := "SELECT id, nombre, username, password FROM usuarios WHERE username=?"
	return getUsuario(sql, username)
}

//GetUsuarios function
func GetUsuarios() Usuarios {
	var usuarios Usuarios
	sql := "SELECT id, nombre, username, password FROM usuarios"
	rows, _ := Query(sql)
	for rows.Next() {
		var usuario Usuario
		rows.Scan(&usuario.ID, &usuario.Nombre, &usuario.Username, &usuario.Password)
		usuarios = append(usuarios, usuario)
	}
	return usuarios
}

func LoginUsuario(username, password string) (*Usuario, error) {
	usuario,_ := GetUsuarioByUsername(username)
	if usuario.Password != password {
		return &Usuario{}, errors.New("Usuario o contraseña no coinciden")
	}
	return usuario, nil
}

//Save method
func (usuario *Usuario) Save() error {
	if usuario.ID == 0 {
		return usuario.insert()
	}
	return usuario.update()
}

func (usuario *Usuario) insert() error {
	sql := "INSERT INTO usuarios(nombre, username, password) VALUES(?,?,?);"
	usuarioID, err := InsertData(sql, usuario.Nombre, usuario.Username, usuario.Password)
	usuario.ID = int(usuarioID)
	return err
}

func (usuario *Usuario) update() error {
	sql := "UPDATE usuarios SET nombre=?, username=?, passwprd=? WHERE id=?"
	_, err := Exec(sql, usuario.Nombre, usuario.Username, usuario.Password, usuario.ID)
	return err
}

//Delete method
func (usuario *Usuario) Delete() error {
	sql := "DELETE FROM usuarios WHERE id=?"
	_, err := Exec(sql, usuario.ID)
	return err
}

func (usuario *Usuario) SetPassword(password string) {
	usuario.Password = password
}


func (usuario *Usuario) CrearCuenta(cuenta Cuenta) {
	
}
