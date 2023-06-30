package pkg

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysql() (*sql.DB, error) {

	var MYSQL_HOST = "cifrado.com.co"
	var MYSQL_USER = "cifrados_drako"
	var MYSQL_PASS = "Desarrollo$123"
	var MYSQL_PORT = 3306
	var MYSQL_DB = "cifrados_dev"

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

	var query2 = "CREATE TABLE IF NOT EXISTS books ( id INT PRIMARY KEY, title VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, literary_genre VARCHAR(255) NOT NULL, created_at TIMESTAMP DEFAULT NOW() );"
	db.Query(query2)

	var query3 = "CREATE TABLE IF NOT EXISTS lend_books ( id INT PRIMARY KEY, user_id INT NOT NULL, book_id INT NOT NULL, return_book TIMESTAMP, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(), FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE, FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE );"
	db.Query(query3)

	fmt.Println("Connexion a base de datos Mysql ok")
	return db, nil
}
