package main

import (
	"dev-cli/internal/env"
	"dev-cli/internal/jsonformat"
	"dev-cli/internal/uuid"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Помилка: будь ласка, вкажіть команду. Наприклад: dev uuid")
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "uuid":
		key := uuid.Generate()
		fmt.Println("New key:", key)
	case "env":
		env.PrintAll()
	case "json":
		if len(os.Args) < 3 {
			fmt.Println("Помилка: вкажіть JSON-рядок. Наприклад: dev json '{\"a\":1}'")
			os.Exit(1)
		}
		formated, err := jsonformat.Format(os.Args[2])
		if err != nil {
			fmt.Println("Помилка: невалідний JSON формат ->", err)
			os.Exit(1)
		}
		fmt.Println(formated)
	default:
		fmt.Printf("Невідома команда: '%s'. Спробуйте 'uuid'.\n", command)
	}
}
