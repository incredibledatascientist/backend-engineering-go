package postgres

// import (
// 	"context"
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"jwt-auth/internal/domain"
// )

// type UserStore struct {
// 	db *sql.DB
// }

// func NewUserStore(db *sql.DB) *UserStore {
// 	return &UserStore{db: db}
// }

// func (p *UserStore) CreateUser(ctx context.Context, u *domain.User) error {
// 	query := `
// 	INSERT INTO users (name, email, password)
// 	VALUES ($1, $2, $3)
// 	RETURNING id, created_at;
// 	`
// 	return p.db.QueryRowContext(
// 		ctx,
// 		query,
// 		u.FirstName,
// 		u.Email,
// 		u.Password,
// 	).Scan(&u.ID, &u.CreatedAt)
// }

// func (p *UserStore) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
// 	query := `
// 	SELECT id, name, email, password, created_at
// 	FROM users
// 	WHERE id=$1;
// 	`
// 	u := &domain.User{}
// 	err := p.db.QueryRowContext(ctx, query, id).Scan(
// 		&u.ID,
// 		&u.FirstName,
// 		&u.Email,
// 		&u.Password,
// 		&u.CreatedAt,
// 	)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, fmt.Errorf("user with id %d not found", id)
// 		}
// 		return nil, err
// 	}
// 	return u, nil
// }

// func (p *UserStore) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
// 	query := `
// 	SELECT id, name, email, password, created_at
// 	FROM users
// 	WHERE email=$1;
// 	`
// 	u := &domain.User{}
// 	err := p.db.QueryRowContext(ctx, query, email).Scan(
// 		&u.ID,
// 		&u.FirstName,
// 		&u.Email,
// 		&u.Password,
// 		&u.CreatedAt,
// 	)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, fmt.Errorf("user with email %s not found", email)
// 		}
// 		return nil, err
// 	}
// 	return u, nil
// }

// func (p *UserStore) GetUsers(ctx context.Context) ([]*domain.User, error) {
// 	query := `
// 	SELECT id, name, email, password, created_at
// 	FROM users;
// 	`
// 	rows, err := p.db.QueryContext(ctx, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []*domain.User
// 	for rows.Next() {
// 		u := &domain.User{}
// 		if err := rows.Scan(
// 			&u.ID,
// 			&u.FirstName,
// 			&u.Email,
// 			&u.Password,
// 			&u.CreatedAt,
// 		); err != nil {
// 			return nil, err
// 		}
// 		users = append(users, u)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// func (p *UserStore) UpdateUser(ctx context.Context, u *domain.User) error {
// 	query := `
// 	UPDATE users
// 	SET name=$1, email=$2, password=$3
// 	WHERE id=$4;
// 	`
// 	result, err := p.db.ExecContext(
// 		ctx,
// 		query,
// 		u.FirstName,
// 		u.Email,
// 		u.Password,
// 		u.ID,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("user with id %d not found", u.ID)
// 	}
// 	return nil
// }

// func (p *UserStore) DeleteUser(ctx context.Context, id int) error {
// 	query := `DELETE FROM users WHERE id=$1`
// 	result, err := p.db.ExecContext(ctx, query, id)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("user with id %d not found", id)
// 	}
// 	return nil
// }

// func (p *UserStore) WithTx(ctx context.Context, fn func(*sql.Tx) error) error {
// 	tx, err := p.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}
// 	if err := fn(tx); err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx err: %v, rollback err: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }
