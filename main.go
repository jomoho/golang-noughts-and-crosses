package main

import (
    "net/http"
    "log"
	"./xo"
)

func main() {
	router := xo.Router()
    log.Fatal(http.ListenAndServe(":8000", router))
}