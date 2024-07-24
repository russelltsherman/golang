package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getIndex \n")
	w.Header().Set("Content-Type", "application/json")
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("createMovie \n")
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("deleteMovies \n")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getMovie \n")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("getMovies \n")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("updateMovie \n")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	movies = append(movies, Movie{ID: "1", Isbn: "2432143", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "5647532", Title: "Movie Two", Director: &Director{Firstname: "John", Lastname: "Doe"}})

	router := mux.NewRouter()
	router.HandleFunc("/", getIndex).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
