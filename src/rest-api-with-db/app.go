package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	model "./model"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	dao "rest-api-with-db/dao"
	"strconv"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/movies", AllMoviesEndpoint).Methods("GET")
	r.HandleFunc("/movies/{id}", FindMovieEndpoint).Methods("GET")
	r.HandleFunc("/movies", CreateMovieEndpoint).Methods("POST")
	r.HandleFunc("/movies", UpdateMovieEndpoint).Methods("PUT")
	r.HandleFunc("/movies", DeleteMovieEndpoint).Methods("DELETE")

	if err:=http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

func AllMoviesEndpoint(w http.ResponseWriter, r *http.Request) {
	movies, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func FindMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pending")
	
}

func CreateMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pending")
	
}

func UpdateMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pending")
	
}

func DeleteMovieEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pending")
	
}