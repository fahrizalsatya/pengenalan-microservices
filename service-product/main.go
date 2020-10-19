package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FahrizalSatya/pengenalan-microservices/service-product/handler"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/add-product", http.HandlerFunc(handler.AddMenuHandler))

	fmt.Println("Server listen on :8000")
	log.Panic(http.ListenAndServe(":8000", router))

}
