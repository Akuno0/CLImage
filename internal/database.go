// Для работы с БД
package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=12345678 dbname=images host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Подключено")

	return db, nil
}

func GetHash(db *sql.DB, fileID string) (string, error) {
	query := "SELECT hash FROM files WHERE id = $" + fileID + ";"
	row := db.QueryRow(query, fileID)

	var hash string
	if err := row.Scan(&hash); err != nil {
		return "", fmt.Errorf("не удалось получить хэш: %v", err)
	}

	return hash, nil
}
