// Основная логика
package internal

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func Run(hash string) error {
	path := fmt.Sprintf("../images/%s.jpg", hash)
	src, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("Ошибка открытия изображения: %v", err)
	}

	dst := imaging.Resize(src, 100, 100, imaging.Lanczos)

	err = imaging.Save(dst, fmt.Sprintf("../images/%s_s.jpg", hash))
	if err != nil {
		log.Fatalf("Ошибка сохраниения изображения: %v", err)
	}

	src1, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("Ошибка открытия изображения: %v", err)
	}
	dst1 := imaging.Resize(src1, 0, 256, imaging.Lanczos)

	err = imaging.Save(dst1, fmt.Sprintf("../images/%s_m.jpg", hash))
	if err != nil {
		log.Fatalf("Ошибка сохраниения изображения: %v", err)
	}
	return nil
}
