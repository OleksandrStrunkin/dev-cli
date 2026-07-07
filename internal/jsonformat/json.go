package jsonformat

import (
	"bytes"
	"encoding/json"
)

func Format(input string) (string, error) {
	inBytes := []byte(input)

	var outBytes bytes.Buffer

	err := json.Indent(&outBytes, inBytes, "", "  ")

	if err != nil {
		return "", err
	}
	return outBytes.String(), nil
}
