package main

import (
	"log"
	"net/http"

	"skyitachi/mesh_agent/controller"
)

func main() {
	http.HandleFunc("/", controller.HelloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
