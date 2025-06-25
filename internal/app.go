// Основная логика
package internal

/*"bytes"
"crypto/sha256"
"fmt"
"image"
"image/jpeg"
"os"*/

import (
	"fmt"
	"os"
)

func Run(hash string) error {
	createDir(hash)

	return nil
}

func createDir(hash string) {
	err := os.Mkdir(hash, 0755)
	if err != nil {
		fmt.Println("Ошибка при создании директории:", err)
		return
	}
}
