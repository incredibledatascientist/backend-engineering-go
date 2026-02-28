package handler

import (
	"encoding/json"
	"jwt-auth/internal/domain"
	"jwt-auth/internal/storage"
	"jwt-auth/internal/utils"
	"net/http"
)

type MovieHandler struct {
	store storage.MovieStorage
}

func NewMovieHandler(store storage.MovieStorage) *MovieHandler {
	return &MovieHandler{store: store}
}

func (m *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.store.GetMovies()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Server error", err.Error())
		return
	}
	utils.Success(w, http.StatusOK, "Movies fetched successfully", movies)
}

func (m *MovieHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request parameter", err.Error())
		return
	}

	movie, err := m.store.GetMovie(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "Movie doesnt exists", err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "Movie fetched successfully", movie)
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	req := domain.MovieRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if req.Title == "" || req.Year == 0 {
		utils.Error(w, http.StatusBadRequest, "Validation error", "Movie title & year are required")
		return
	}
	movie, err := domain.NewMovie(req.Title, req.Year)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Movie creation failed", err.Error())
		return
	}

	if err := m.store.CreateMovie(movie); err != nil {
		utils.Error(w, http.StatusInternalServerError, "Create movie err", err.Error())
		return
	}
	utils.Success(w, http.StatusOK, "Movie created successfully", movie)
}

func (m *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request parameter", err.Error())
		return
	}

	err = m.store.DeleteMovie(id)
	if err != nil {
		utils.Error(w, http.StatusNotFound, "Movie not present", err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "Movie deleted successfully", nil)
}

func (m *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id, err := ParseIntParam(r, "id")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	movie := &domain.Movie{}
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		utils.Error(w, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	movie.ID = id

	movie, err = m.store.UpdateMovie(movie)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, "Failed to update movie", err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "Movie updated successfully", movie)
}
