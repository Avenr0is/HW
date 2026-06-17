package file

import (
	"os"
	"strings"
)

func ReadFile(name string) (string, error) {
	readed, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(readed), nil
}

func WriteFile(name string, data []byte) error {
	err := os.WriteFile(name, data, 0664)
	if err != nil {
		return err
	}
	return nil
}

func FormatCheck(name string) bool {
	return strings.Contains(name, ".json")
}