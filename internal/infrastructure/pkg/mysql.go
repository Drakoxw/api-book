package pkg

import (
	"api-book/internal/infrastructure/utils"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() (*sql.DB, error) {

	var MYSQL_HOST = utils.GetEnvVariable("MYSQL_HOST")
	var MYSQL_USER = os.Getenv("MYSQL_USER")
	var MYSQL_PASS = "Desarrollo$123"
	var MYSQL_PORT = os.Getenv("MYSQL_PORT")
	var MYSQL_DB = os.Getenv("MYSQL_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", MYSQL_USER, MYSQL_PASS, MYSQL_HOST, MYSQL_PORT, MYSQL_DB)
	log.Println(dsn)

	// Establecer la conexión a la base de datos
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	// Verifica la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("No hubo respuesta de la base de datos: %v", err)
	}

	var query = "CREATE TABLE if not exists users (id INT PRIMARY KEY, username VARCHAR ( 50 ) UNIQUE NOT NULL, password VARCHAR ( 150 ) NOT NULL, email VARCHAR ( 255 ) UNIQUE NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NULL)"
	db.Query(query)

	fmt.Println("Connexion a base de datos Mysql ok")
	return db, nil
}
