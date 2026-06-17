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
   ReadBin(&file, data)
}

func ReadBin(file *os.File, data []byte)(strFlData string){
    
	readed, err := file.Read(data)

	if err != nil{
		return 
	}
   
	strFlData = string(readed)

	return strFlData
}

func CreatedBinsFl()(*os.File){
     
	file, err := os.Create("binsfl.json")

	if err != nil{

		return  nil
	}
   
	return file
}




