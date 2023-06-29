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

	fmt.Println("Connexion a base de datos Postgres ok")
	return db, nil
}
