package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

type logWriter struct {}

func main(){
	resp, err := http.Get("http://www.google.com");
	if(err != nil){
		fmt.Println("Error", err);
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error){
	fmt.Println(string(bs))
	fmt.Println("Just wrote out this many bites :", len(bs))
	return len(bs), nil
}