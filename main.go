package main

import (
	"log"
	"net/http"

	service "ascii/service"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", service.MainPageHandler)
	http.HandleFunc("/ascii-art/", service.PostPageHandler)
	log.Println("Server is listening")
	log.Println(http.ListenAndServe(port, nil))
}
