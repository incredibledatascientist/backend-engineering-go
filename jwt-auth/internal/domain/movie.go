package domain

import (
	"jwt-auth/internal/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID        string    `json:"id"`
	Year      int       `json:"year"`
	Title     string    `json:"title"`
	Director  *Director `json:"director"`
	CreatedAt time.Time `json:"created_at"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Movie Handler, Later use dependency injection.
func GetMovies(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{"movies": []Movie{}}
	utils.Success(w, http.StatusOK, "Movies fetched successfully", response)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := map[string]any{"params": params}
	utils.Success(w, http.StatusOK, "Movie fetched successfully", response)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{"movies": Movie{}}
	utils.Success(w, http.StatusOK, "Movie created successfully", response)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{"movies": Movie{}}
	utils.Success(w, http.StatusOK, "Movie updated successfully", response)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	response := map[string]any{"params": params}
	utils.Success(w, http.StatusOK, "Movie deleted successfully", response)
}
