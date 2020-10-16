package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rezaahmadk/dts-microservice/auth-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/admin-auth", http.HandlerFunc(handler.ValidateAuth))

	fmt.Printf("Auth service listen on :8001")
	log.Panic(http.ListenAndServe(":8001", router))
}
