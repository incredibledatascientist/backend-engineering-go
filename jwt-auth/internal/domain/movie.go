package domain

import (
	"time"
)

type Movie struct {
	ID        int       `json:"id"`
	Year      int       `json:"year"`
	Title     string    `json:"title"`
	Director  *Director `json:"director"`
	CreatedAt time.Time `json:"created_at"`
}

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type MovieRequest struct {
	Year  int    `json:"year"`
	Title string `json:"title"`
	// Director  *Director `json:"director"`
}

func NewMovie(title string, year int) (*Movie, error) {
	movie := &Movie{
		Title: title,
		Year:  year,
	}
	return movie, nil
}
