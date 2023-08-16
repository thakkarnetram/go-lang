package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {	
	fmt.Println("Modules in go")
	hello()
	r:=mux.NewRouter()
	r.HandleFunc("/",getHome).Methods("GET")

	// making server
	log.Fatal(http.ListenAndServe(":4000",r))
}

func hello(){
	fmt.Println("Hello Users")
}

func getHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h2> Hello <3 </h2>"))
}