// Для работы с БД
package internal

import (
	"database/sql"
	"fmt"
	"log"
)

func db() {
	connStr := "user=user password=123 dbname=db host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Подключено")
}
