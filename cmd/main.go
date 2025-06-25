// Точка входа
package main

import (
	//"context"
	"fmt"
	"os"

	"github.com/Akuno0/CLImage/internal"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("go run main.go <file_id>")
		os.Exit(1)
	}

	fileID := os.Args[1]

	// Инициализация подключения к Postgres
	db, err := internal.InitDB()
	if err != nil {
		fmt.Printf("Ошибка инициализации: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Получение hash
	hash, err := internal.GetHash(db, fileID)
	if err != nil {
		fmt.Printf("Ошибка получения хэша: %v\n", err)
		os.Exit(1)
	}

	// Генерация
	err = internal.Run(hash)
	if err != nil {
		fmt.Printf("Ошибка генерации: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Готво.")
}
