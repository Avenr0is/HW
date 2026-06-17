package storage

import (
	"encoding/json"
	"project/bins"
	"os"
)


func BinSaved(file os.File){

	var myBin bins.Bin
	data, err := json.Marshal(myBin)

	if err != nil{
		return
	}

	file = *CreatedBinsFl()
	_, err = file.Write(data)
	
	if err != nil{
		return
	}
   ReadBin(&myBin, data)
}
func ReadBin(myBin *bins.Bin, data []byte) string {

    err := json.Unmarshal(data, myBin)
    if err != nil {
        return "Ошибка: не удалось дессериализовать файл"
    }
    
    return string(data)
}

func CreatedBinsFl()(*os.File){
     
	file, err := os.Create("binsfl.json")

	if err != nil{

		return  nil
	}
   
	return file
}




