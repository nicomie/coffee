package db

import (
	"database-interface/db/sqlc"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
}

func ListCoffees(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	coffees, err := Queries.ListCoffees(CTX)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(coffees)
}

func GetCoffee(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
	}

	fetchedCoffee, err := Queries.GetCoffee(CTX, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(fetchedCoffee)

}
func NewCoffee(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var coffee sqlc.Coffee

	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		fmt.Println(err)
	}

	path := fmt.Sprintf("storage/images/%s", coffee.Name)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.Write([]byte(coffee.ImageSrc))

	insertedCoffee, err := Queries.CreateCoffee(CTX, sqlc.CreateCoffeeParams{
		Name:     coffee.Name,
		Flavor:   coffee.Flavor,
		Acidity:  coffee.Acidity,
		ImageSrc: coffee.ImageSrc,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(insertedCoffee)
}

func DeleteCoffee(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	params := mux.Vars(r)
	idString := params["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
	}

	err = Queries.DeleteCoffee(CTX, int64(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func UpdateCoffee(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var coffee sqlc.Coffee
	var id int

	err := json.NewDecoder(r.Body).Decode(&coffee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	idString := params["id"]
	id, err = strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
	}

	author, err := Queries.UpdateCoffee(CTX, sqlc.UpdateCoffeeParams{
		ID:       int64(id),
		Name:     coffee.Name,
		Flavor:   coffee.Flavor,
		Acidity:  coffee.Acidity,
		ImageSrc: coffee.ImageSrc,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return

	}

	json.NewEncoder(w).Encode(author)
}
