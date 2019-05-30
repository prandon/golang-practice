package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Person struct {
	ID 			string `json:"id,omitempty"`
	FirstName 	string `json:"firstname,omitempty"`
	LastName 	string `json:"lastname,omitempty"`
	Address 	*Address `json:"address,omitempty"`
}

type Address struct {
	City 	string `json:"city,omitempty"`
	State 	string `json:"state,omitempty"`
}

var people []Person
 	
func main() {

	people = append(people, Person{ID: "1", FirstName: "Pranjal", LastName: "Nartam", Address: &Address{City: "Pandharkawada", State: "Maharashtra"}})
	people = append(people, Person{ID: "2", FirstName: "Rasika", LastName: "Warang", Address: &Address{City: "Humras", State: "Maharashtra"}})
	people = append(people, Person{ID: "3", FirstName: "Siddhant", LastName: "Nartam", Address: &Address{City: "Pandharkawada", State: "Maharashtra"}})
	people = append(people, Person{ID: "4", FirstName: "Subhash", LastName: "Nartam", Address: &Address{City: "Pandharkawada", State: "Maharashtra"}})

	router := mux.NewRouter()

	router.HandleFunc("/people", GetPeople).Methods("GET");	
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET");
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST");
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE");

	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = param["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	for index, item := range people {
		if item.ID == param["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}