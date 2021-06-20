package main

import (
	"fmt"
	"os"
	"io"
)


func main(){
	fileName := os.Args[1];
	file, err := os.Open(fileName);
	if(err != nil){
		fmt.Println("error cocured : ", err);
	}

	io.Copy(os.Stdout, file)
}
