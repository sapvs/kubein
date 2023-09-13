package main

import (
	"log"
	"net/http"

	"github.com/sapvs/kubein"
)

func main() {
	http.HandleFunc("/", kubein.Root)
	log.Println("starting server")
	go http.ListenAndServe(":8080", http.DefaultServeMux)
	<-kubein.Shutdown()
	log.Println("donee")
}
