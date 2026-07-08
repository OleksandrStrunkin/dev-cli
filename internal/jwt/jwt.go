package jwt

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func DecodePayload(token string) (string, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return "", fmt.Errorf("Iinvalid JWT format (must have 3 parts)")
	}

	payloadSegment := parts[1]

	decodeBytes, err := base64.RawURLEncoding.DecodeString(payloadSegment)

	if err != nil {
		return "", fmt.Errorf("decoding error Base64: %v", err)
	}

	return string(decodeBytes), nil
}
