package main

import (
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

type Result struct {
	Status string
}

func main(){
	http.HandleFunc("/", home) 
	http.ListenAndServe(":9094", nil)


}

func home(w http.ResponseWriter, r*http.Request){

	result := Result{Status: "worked"}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error on json")
	}
	fmt.Fprintf(w, string(jsonResult))
}