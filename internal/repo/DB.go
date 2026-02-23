package repo

import (
	cfg "auth/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func StartBD() (*sql.DB, error) {

	cb, err := cfg.LoadConfigBD()

	if err != nil {
		log.Fatal("Error Loading Configurate BD")
		return nil, fmt.Errorf("Error Loading Configurate BD:", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cb.BD_USER, cb.BD_PASSWORD, cb.BD_HOST, cb.BD_PORT, cb.DB_NAME)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Успешное подключение!")

	return db, nil

}
