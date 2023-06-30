package pkg

import (
	"api-book/internal/infrastructure/utils"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func InitPostgres() (*sql.DB, error) {

	var POSTGRES_HOST = utils.GetEnvVariable("POSTGRES_HOST")
	var POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	var POSTGRES_USER = os.Getenv("POSTGRES_USER")
	var POSTGRES_PASS = os.Getenv("POSTGRES_PASS")
	var POSTGRES_DB = os.Getenv("POSTGRES_DB")

	dns := fmt.Sprintf("postgres://%s:%s@%s:%v/%s", POSTGRES_USER, POSTGRES_PASS, POSTGRES_HOST, POSTGRES_PORT, POSTGRES_DB)

	db, err := sql.Open("postgres", dns)

	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("No hubo respuesta de la base de datos: %v", err)
	}

	var query = "CREATE TABLE if not exists users ( id serial PRIMARY KEY, username VARCHAR ( 50 ) UNIQUE NOT NULL, password VARCHAR ( 150 ) NOT NULL, email VARCHAR ( 255 ) UNIQUE NOT NULL, created_at TIMESTAMP NOT NULL, updated_at TIMESTAMP NULL)"
	db.Query(query)

	var query2 = "CREATE TABLE IF NOT EXISTS books ( id SERIAL PRIMARY KEY, title VARCHAR(255) NOT NULL, author VARCHAR(255) NOT NULL, literary_genre VARCHAR(255) NOT NULL, created_at TIMESTAMP DEFAULT NOW() );"
	db.Query(query2)

	var query3 = "CREATE TABLE IF NOT EXISTS lend_books ( id SERIAL PRIMARY KEY, user_id INT NOT NULL, book_id INT NOT NULL, return_book TIMESTAMP, created_at TIMESTAMP DEFAULT NOW(), updated_at TIMESTAMP DEFAULT NOW(), FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE, FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE );"
	db.Query(query3)

	// fmt.Println("Connexion a base de datos Postgres ok")
	return db, nil
}
