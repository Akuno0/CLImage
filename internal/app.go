// Основная логика
package internal

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func Run(fileID string) error {
	file, err := os.Open(fileID)
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

	fmt.Println("Хэш изображения:", hashString)
	return nil
}
