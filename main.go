package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Director *Director `json:"director"`
}
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	// initialisation of mux for routing
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "001", Name: "RRR", Director: &Director{Firstname: "RAJA", Lastname: "MOLI"}})
	movies = append(movies, Movie{ID: "002", Name: "BB", Director: &Director{Firstname: "RAJAN", Lastname: "MOOLII"}})
        // common endpoints in every application 
	r.HandleFunc("/movies", getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getmovie).Methods("GET")
	r.HandleFunc("/movies", creatmovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deletemovies).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
func getmovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)

}
func deletemovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for index, item := range movies {
		if params["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(movies)

}
func getmovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range movies {
		if params["id"] == item.ID {

			json.NewEncoder(w).Encode(item)
			return

		}
	}

}
func creatmovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa((rand.Intn(1000000)))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)

}
func updatemovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parmes := mux.Vars(req)
	for index, item := range movies {
		if parmes["id"] == item.ID {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = strconv.Itoa((rand.Intn(1000000)))
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)

		}

	}
}
