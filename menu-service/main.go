package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rezaahmadk/dts-microservice/menu-service/handler"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/add-menu", http.HandlerFunc(handler.AddMenu))

	fmt.Println("Menu Service Listen on PORT 8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
