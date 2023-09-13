package kubein

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(struct{ Name string }{Name: "test"})
	log.Println("retruned respose")
}

func Shutdown() <-chan struct{} {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGQUIT)

	doneC := make(chan struct{})

	go func(sigChannel chan os.Signal, doneChannel chan struct{}) {
		<-sigChannel
		log.Println("receved term bye")
		doneChannel <- struct{}{}
	}(sig, doneC)

	return doneC
}
