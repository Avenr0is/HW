package file

import (
	"errors"
	"os"
	"fmt"
	"strings"
)

func ReadFile(name string) {

	readed, err := os.ReadFile(name)
	if err != nil{

	 errors.New("Ошибка чтения файла")
	 return 

	} 

	fmt.Println(readed)
}

func WriteFile(name string, data []byte ){

	 err := os.WriteFile(name, data, 0664)

	if err != nil{

		errors.New("Ошибка записи")

	}


}

func FormatCheck(name string)(format bool){

	format = strings.Contains(name,".json")

 return format
}
