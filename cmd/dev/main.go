package main

import (
	"fmt"
	"os"
	"dev-cli/internal/uuid"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Помилка: будь ласка, вкажіть команду. Наприклад: dev uuid")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "uuid":
		key:= uuid.Generate();
		fmt.Println("New key:", key)

	default:
		fmt.Printf("Невідома команда: '%s'. Спробуйте 'uuid'.\n", command)
	}
}
