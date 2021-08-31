package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main(){

	helloHandler := func(w http.ResponseWriter, req *http.Request){
		w.Header().Set("Access-Control-Allow-Origin", "*")
		

		fmt.Println("received:",req.URL)
		if(strings.Contains(req.URL.String(),"ping")){
			io.WriteString(w,"Pong")
		}else{
			io.WriteString(w,"Ping")
		}
	}
	
	http.HandleFunc("/", helloHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))

}