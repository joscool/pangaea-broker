package main

import (
	"log"
	"net/http"

	"github.com/custom-broker/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/subscribe/{topic}", handlers.HandleSubsription)
	myRouter.HandleFunc("/publish/{topic}", handlers.HandlePublishing)
	myRouter.HandleFunc("/event", handlers.HandleEevent)
	log.Println("server started on port 8000")
	log.Fatal(http.ListenAndServe(":8000", myRouter))

}
