package file

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) {

	readed, err := os.ReadFile(name)
	if err != nil {
		return
	}

	fmt.Println(readed)
}

func WriteFile(name string, data []byte) {

	err := os.WriteFile(name, data, 0664)

	if err != nil {

		return

	}

}

func FormatCheck(name string) (format bool) {

	format = strings.Contains(name, ".json")

	return format
}
