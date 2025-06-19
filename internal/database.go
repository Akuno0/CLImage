// Для работы с БД
package internal

import (
	//"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=12345678 dbname=music host=localhost port=5432 sslmode=disable"
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

func GetHash(fileID string) {
	/*file, err := os.Open(fileID)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Декодирование изображения в объект image.Image
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Получение байт изображения
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		panic(err)
	}
	imageData := buf.Bytes()

	// Вычисление хэша SHA256
	hash := sha256.New()
	hash.Write(imageData)
	hashBytes := hash.Sum(nil)

	hashString := fmt.Sprintf("%x", hashBytes)

	fmt.Println("Хэш изображения:", hashString)*/
	return
}
