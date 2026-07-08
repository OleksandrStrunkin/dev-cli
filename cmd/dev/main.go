package main

import (
	"dev-cli/internal/env"
	"dev-cli/internal/hash"
	"dev-cli/internal/jsonformat"
	"dev-cli/internal/jwt"
	"dev-cli/internal/server"
	"dev-cli/internal/uuid"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: please provide a command. Example: dev uuid")
		fmt.Println("Available commands: uuid, env, json, hash, serve, jwt, info")
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
			fmt.Println("Error: please provide a JSON string. Example: dev json '{\"a\":1}'")
			os.Exit(1)
		}
		formatted, err := jsonformat.Format(os.Args[2])
		if err != nil {
			fmt.Println("Error: invalid JSON format ->", err)
			os.Exit(1)
		}
		fmt.Println(formatted)
        
	case "hash":
		if len(os.Args) < 4 {
			fmt.Println("Error: insufficient arguments.")
			fmt.Println("Usage: dev hash <algorithm> <text>")
			fmt.Println("Algorithms: sha256, md5")
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
			fmt.Printf("Error: unknown algorithm '%s'. Available: sha256, md5\n", algorithm)
			os.Exit(1)
		}

	case "serve":
		port := ":8080"
		server.Start(port)

	case "jwt":
		if len(os.Args) < 3 {
			fmt.Println("Error: please provide a JWT token. Example: dev jwt <token>")
			os.Exit(1)
		}

		token := os.Args[2]

		rawJSON, err := jwt.DecodePayload(token)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		formatted, err := jsonformat.Format(rawJSON)
		if err != nil {
			fmt.Println("Error formatting decoded JSON:", err)
			os.Exit(1)
		}

		fmt.Println(formatted)
        
	case "info":
		fmt.Println("Available commands: uuid, env, json, hash, serve, jwt, info")
        
	default:
		fmt.Printf("Error: unknown command '%s'. Try 'dev info' to see available commands.\n", command)
		os.Exit(1)
	}
}
