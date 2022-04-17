package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MovieDetail struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Genre       string   `json:"genre"`
	ReleaseYear int      `json:"releaseYear"`
	RunningTime int      `json:"runningTime"`
	Director    string   `json:"director"`
	Stars       []string `json:"actors"`
}

var cache = map[int]MovieDetail{}

func initCache() {
	defaultMovieDetails := []MovieDetail{
		{
			ID:          0,
			Title:       "The Shawshank Redemption",
			Genre:       "Drama",
			ReleaseYear: 1994,
			RunningTime: 142,
			Director:    "Frank Darabont",
			Stars:       []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton"},
		},
		{
			ID:          1,
			Title:       "The Godfather",
			Genre:       "Drama, Thriller",
			ReleaseYear: 1972,
			RunningTime: 175,
			Director:    "Francis Ford Coppola",
			Stars:       []string{"Marlon Brando", "Al Pacino", "James Caan"},
		},
		{
			ID:          2,
			Title:       "The Dark Knight",
			Genre:       "Action, Drama",
			ReleaseYear: 2008,
			RunningTime: 152,
			Director:    "Christopher Nolan",
			Stars:       []string{"Christian Bale", "Heath Ledger", "Aaron Eckhart"},
		},
	}

	for _, detail := range defaultMovieDetails {
		cache[detail.ID] = detail
	}
}

func GetMovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	detail, ok := cache[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No such detail (ID: ", id, ")")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(detail)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

}

func CreateMovieDetailHandler(w http.ResponseWriter, r *http.Request) {
	var movieDetail MovieDetail
	err := json.NewDecoder(r.Body).Decode(&movieDetail)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	cache[movieDetail.ID] = movieDetail

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(movieDetail)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

}

func NewHandler() http.Handler {
	initCache()
	router := mux.NewRouter()

	router.HandleFunc("/details/{id:[0-9]+}", GetMovieDetailHandler).Methods("GET")
	router.HandleFunc("/details", CreateMovieDetailHandler).Methods("POST")
	return router
}
func main() {
	handler := NewHandler()
	log.Println("Movie Detail Server Listening 8081...")
	log.Fatal(http.ListenAndServe(":8081", handler))
}
