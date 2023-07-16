package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// NewDB создает и возвращает новый экземпляр *sql.DB
func NewDB() (*sql.DB, error) {
	// Параметры подключения к базе данных PostgreSQL
	dbHost := "localhost"
	dbPort := 5432
	dbUser := "mustaz"
	dbPassword := "qwe123"
	dbName := "mydatabase"

	// Формирование строки подключения
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Установка соединения с базой данных
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Проверка соединения с базой данных
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
