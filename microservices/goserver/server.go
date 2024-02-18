package main

import (
	"fmt"
	"goserver/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	fmt.Print("Server ready")

	fmt.Printf("The go server is running on port 3010 ...\n")
	log.Fatal(http.ListenAndServe(":3010", router))
}
