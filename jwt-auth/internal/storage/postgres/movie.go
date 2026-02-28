package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-auth/internal/domain"
)

func (p *PostgresStore) CreateMovieTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS movie (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			year int NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		`

	_, err := p.db.Exec(query)
	return err
}

func (p *PostgresStore) CreateMovie(movie *domain.Movie) error {
	query := `
		INSERT INTO movie (title, year)
		VALUES ($1, $2)
		RETURNING id
	`

	return p.db.QueryRow(query, movie.Title, movie.Year).Scan(&movie.ID)
}

// QueryRow will give err if no rec found but Query will retun nil
func (p *PostgresStore) GetMovie(id int) (*domain.Movie, error) {
	query := `SELECT id, title, year, created_at FROM movie WHERE id=$1`

	movie := &domain.Movie{}

	err := p.db.QueryRow(query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Year,
		&movie.CreatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("Movie not found")
	}

	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (p *PostgresStore) DeleteMovie(id int) error {
	_, err := p.GetMovie(id)
	if err != nil {
		return err
	}

	query := `DELETE FROM movie WHERE id=$1`
	_, err = p.db.Query(query, id)
	return err
}

func (p *PostgresStore) GetMovies() ([]*domain.Movie, error) {
	query := `SELECT id, title, year, created_at FROM movie`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var movies []*domain.Movie
	for rows.Next() {
		movie := &domain.Movie{}
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			// &movie.Director,
			&movie.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil

}

func (p *PostgresStore) UpdateMovie(movie *domain.Movie) (*domain.Movie, error) {
	query := `
		UPDATE movie  SET title = $1, year = $2 WHERE id = $3
		RETURNING id, title, year, created_at
	`

	updated := &domain.Movie{}

	err := p.db.QueryRow(
		query,
		movie.Title,
		movie.Year,
		movie.ID,
	).Scan(
		&updated.ID,
		&updated.Title,
		&updated.Year,
		&updated.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return updated, nil
}
