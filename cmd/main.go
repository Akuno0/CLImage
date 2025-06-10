// Точка входа
package main

import (
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

	if err := internal.Run(fileID); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Success")
}
