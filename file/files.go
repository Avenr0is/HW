package file

import (
	"os"
	"strings"
	
)

func ReadFile(name string)(strReded string) {

	readed, err := os.ReadFile(name)
	if err != nil {
		return "Не удалось прочитать файл"
	}

	strReded = string(readed)
	return 
}

func WriteFile(name string, data []byte){

	err := os.WriteFile(name, data, 0664)

	if err != nil {
		return
	}
	
   return 
}

func FormatCheck(name string) (format bool) {

	format = strings.Contains(name, ".json")

	return format
}
