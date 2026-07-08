package main

import (
	"dev-cli/internal/env"
	"dev-cli/internal/hash"
	"dev-cli/internal/jsonformat"
	"dev-cli/internal/server"
	"dev-cli/internal/uuid"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Помилка: будь ласка, вкажіть команду. Наприклад: dev uuid")
		fmt.Println("Доступні команди: uuid, env, json, info")
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
	case "hash":
		if len(os.Args) < 4 {
			fmt.Println("Помилка: недостатньо аргументів.")
			fmt.Println("Використання: dev hash <алгоритм> <текст>")
			fmt.Println("Алгоритми: sha256, md5")
			os.Exit(1)
		}
		algorithm := os.Args[2]
		textToHash := os.Args[3]

		switch algorithm {
		case "sha256":
			result := hash.SHA256(textToHash)
			fmt.Println(result)
		case "md5":
			result := hash.MD5(textToHash)
			fmt.Println(result)
		default:
			fmt.Printf("Помилка: невідомий алгоритм '%s'. Доступні: sha256, md5\n", algorithm)
			os.Exit(1)
		}

	case "serve":
		port:= ":8080"
		server.Start((port))
	case "info":
		fmt.Println("All command: uuid, env, json")
	default:
		fmt.Printf("Невідома команда: '%s'. Спробуйте 'uuid'.\n", command)
	}
}
