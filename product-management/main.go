package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello, World!")
	})

	err := http.ListenAndServe(":3004", nil)

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
