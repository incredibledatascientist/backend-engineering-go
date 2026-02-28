package postgres

import (
	"jwt-auth/internal/domain"
)

func (p *PostgresStore) CreateMovie(movie *domain.Movie) error {
	query := `
		INSERT INTO movie (title, year)
		VALUES ($1, $2)
		RETURNING id
	`

	return p.db.QueryRow(query, movie.Title, movie.Year).Scan(&movie.ID)
}
