package main

import (
	"database-interface/db"

	store "database-interface/storage"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadConfig()
	store.InitGoogle()

	db.Run()
	fmt.Println("GO ORM tuturoail")
	handleRequests()

}
func helloWord(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", helloWord).Methods("GET")

	myRouter.HandleFunc("/coffee/google/url", store.GetSignedURL).Methods("GET")
	myRouter.HandleFunc("/coffee", db.ListCoffees).Methods("GET")
	myRouter.HandleFunc("/coffee/{id}", db.GetCoffee).Methods("GET")
	myRouter.HandleFunc("/coffee", db.NewCoffee).Methods("POST")
	myRouter.HandleFunc("/coffee/{id}", db.DeleteCoffee).Methods("DELETE")
	myRouter.HandleFunc("/coffee/{id}", db.UpdateCoffee).Methods("PUT")

	log.Fatal(http.ListenAndServe("127.0.0.1:8081", myRouter))

	db.DB.Close()

}
