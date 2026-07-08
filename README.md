# dev-cli

A small Go CLI tool for developers. It contains useful commands for UUID generation, JSON formatting, hashing, starting a simple HTTP server, decoding JWT payloads, and showing system information.

## Features

- Generate UUID
- Print environment variables
- Format JSON strings
- Hash text using SHA256 or MD5
- Run a simple HTTP server
- Decode JWT payloads
- Show system information (Ubuntu only)

## Requirements

- Go 1.26+

## Install

```bash
go build -o dev ./cmd/dev
mv dev /usr/local/bin/
```

> After installing to a directory in your `PATH`, run the tool as `dev`.

## Usage

```bash
dev <command> [arguments]
```

### Available commands

- `uuid` — generate a new UUID
- `env` — print all environment variables
- `json` — format a JSON string
- `hash` — compute a hash of text
- `serve` — start an HTTP server on port 8080
- `jwt` — decode a JWT payload
- `sys` — show system information (Ubuntu only)
- `info` — show available commands

### Examples

Generate a UUID:

```bash
./dev uuid
```

Print environment variables:

```bash
./dev env
```

Format JSON:

```bash
./dev json '{"a":1, "b":2}'
```

Hash text:

```bash
./dev hash sha256 Hello
./dev hash md5 Hello
```

Start the server:

```bash
./dev serve
```

Decode a JWT:

```bash
./dev jwt <token>
```

Show system information (Ubuntu only):

```bash
./dev sys
```

## Project structure

- `cmd/dev` — CLI entry point
- `internal` — command and utility implementations

## License

Open source, use at your own discretion.
