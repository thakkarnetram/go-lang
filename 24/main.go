package main

import (
	"log"
	"net/http"

	"github.com/thakkarnetram/modules/router"
)

func main() {
	r:=router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
}
