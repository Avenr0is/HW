package main

import (
	  "fmt"
	  "project/bins"
      "time"
)

func main() {

   
created()

}


func created() {

	  bin := bins.Bin{

		Id:        "0",
		Private:   true,
	 	CreatedAt: time.Now(),
		Name:      "nonname",
	 }

	 fmt.Println(bin)
	
}
