package models

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
	CreateConnection()
	CreateTables()
}

//CreateConnection method
func CreateConnection() {

	if GetConnection() != nil {
		return
	}

	if connection, err := sql.Open("sqlite3", "./banco.db"); err != nil {
		panic(err)
	} else {
		db = connection
	}
}

//CreateTables method
func CreateTables() {
	createTable("clientes", clienteSchemeSQLITE)
	createTable("tipos_cuenta", tipoCuentaSchemeSQLITE)
	createTable("cuentas", cuentaSchemeSQLITE)
	createTable("empleados", empleadosSchemeSQLITE)
	createTable("tipos_transaccion", tipoTransaccionSchemeSQLITE)
	createTable("transacciones", transaccionSchemeSQLITE)
	createTable("tarjetas", tarjetaSchemeSQLITE)
	
}

/*func createTable(tableName, scheme string) {
	if !existsTable(tableName) {
		Exec(scheme)
	} else {
		truncateTable(tableName)
	}
}*/

func createTable(tableName, scheme string) {
	Exec(scheme)	
}

func truncateTable(tableName string) {
	sql := fmt.Sprintf("DELETE FROM %s", tableName)
	Exec(sql)
}

func existsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, _ := Query(sql)
	return rows.Next()
}

//Exec method
func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	/*if err != nil && !debug {
		log.Println(err)
	}*/
	if err != nil {
		log.Println(err)
	}
	return result, err
}

//Query method
func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	/*if err != nil && !debug {
		log.Println(err)
	}*/
	if err != nil{
		log.Println(err)
	}
	return rows, err
}

//InsertData method
func InsertData(query string, args ...interface{}) (int64, error) {
	result, err := Exec(query, args...)
	if err != nil {
		return int64(0), err
	}
	id, err := result.LastInsertId()
	return id, err
}

//GetConnection method
func GetConnection() *sql.DB {
	return db
}

//Ping method
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//CloseConnection method
func CloseConnection() {
	db.Close()
}
 