package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"project/bins"
)

func BinSaved(bin bins.Bin) {

	data, err := json.Marshal(bin)

	if err != nil {

		errors.New("Ошибка: сереиализация файла была прервана")
	}

	err = os.WriteFile("storage/bins.json", data, 0644)

}

func ReadBin() {

	reded, err := os.ReadFile("storage/bins.json")

	if err != nil {

		errors.New("Ошибка чтения файла")
	}

	fmt.Println(string(reded))
}
