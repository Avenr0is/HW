package storage

import (
	"encoding/json"
	"os"
	"project/bins"
)

func SaveBinsToFile(binsList []bins.Bin, filename string) error {
	data, err := json.Marshal(binsList)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0664)
}

func ReadBinsFromFile(filename string) ([]bins.Bin, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var list []bins.Bin
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}


func CreatedBinsFl() *os.File {
	file, err := os.Create("binsfl.json")
	if err != nil {
		return nil
	}
	return file
}