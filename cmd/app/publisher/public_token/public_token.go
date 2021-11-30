package public_token

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var (
	filePath = os.Getenv("PUBLIC_TOKEN_PEM_PATH")
)

func Generate() (string, error) {
	pem, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Invalid file's path: %w", err)
	}

	keyBase64 := strings.TrimLeft(string(pem), "-----BEGIN PUBLIC KEY-----")
	keyBase64 = strings.TrimRight(keyBase64, "-----END PUBLIC KEY-----")

	keyBinary, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return "", fmt.Errorf("Invalid file's value: %w", err)
	}

	return base64.StdEncoding.EncodeToString(keyBinary), nil
}
