package server

import (
	"jwt-auth/internal/handler"
	"jwt-auth/internal/middleware"
	"jwt-auth/internal/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func (s *HTTPServer) defaultHandler(w http.ResponseWriter, r *http.Request) {
	utils.Error(w, http.StatusNotFound, "The requested endpoint does not exist", nil)
}

func (s *HTTPServer) timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	response := map[string]string{"time": currentTime}
	utils.Success(w, http.StatusOK, "Current time fetched successfully", response)
}

func (s *HTTPServer) healthHandler(w http.ResponseWriter, r *http.Request) {
	health := map[string]string{"status": "ok"}
	utils.Success(w, http.StatusOK, "Service is healthy", health)
}

func (s *HTTPServer) loginHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		utils.Error(w, http.StatusBadRequest, "Form parse error", nil)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	resp := map[string]string{"username": username, "password": password}
	utils.Success(w, http.StatusOK, "Post request successful...!", resp)
}

func (s *HTTPServer) movieRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id, _ := handler.ParseIntParam(r, "id")
		if id != 0 {
			s.movieHandler.GetMovie(w, r)
			return
		} else {
			s.movieHandler.GetMovies(w, r)
			return
		}
	}
	if r.Method == http.MethodPost {
		s.movieHandler.CreateMovie(w, r)
		return
	}
	if r.Method == http.MethodPut {
		s.movieHandler.UpdateMovie(w, r)
		return

	}
	if r.Method == http.MethodDelete {
		s.movieHandler.DeleteMovie(w, r)
		return
	}

	utils.Error(w, http.StatusMethodNotAllowed, "Method not support...!", nil)
}

func (s *HTTPServer) Routes() http.Handler {
	r := mux.NewRouter()
	r.Use(middleware.Logging)

	// staticPath := filepath.Join("internal", "static")
	// fileServer := http.FileServer(http.Dir(staticPath))
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer)).Methods(http.MethodGet)
	// r.HandleFunc("/login", s.loginHandler).Methods(http.MethodPost)

	r.HandleFunc("/time", s.timeHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", s.healthHandler).Methods(http.MethodGet)
	// r.NotFoundHandler = http.HandlerFunc(s.defaultHandler)

	// User Routes
	// r.HandleFunc("/users/signup", h.SignupHandler).Methods(http.MethodPost)
	// r.HandleFunc("/users/login", h.LoginHandler).Methods(http.MethodPost)
	// r.HandleFunc("/users", h.GetUsersHandler).Methods(http.MethodGet)
	// r.HandleFunc("/users/{user_id}", h.GetUserHandler).Methods(http.MethodGet)

	// Movie Routes
	r.HandleFunc("/movies", s.movieRouteHandler)
	r.HandleFunc("/movies/{id}", s.movieRouteHandler)
	// r.HandleFunc("/movies/{id}", s.movieHandler.GetMovie).Methods(http.MethodGet)
	// r.HandleFunc("/movies", s.movieHandler.CreateMovie).Methods(http.MethodPost)
	// r.HandleFunc("/movies/{id}", s.movieHandler.UpdateMovie).Methods(http.MethodPut)
	// r.HandleFunc("/movies/{id}", s.movieHandler.DeleteMovie).Methods(http.MethodDelete)

	return r
}
