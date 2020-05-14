package main

import (
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ola mundo"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Panicln("$PORT not set")
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panicln(err)
	}
}
