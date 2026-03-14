package postgres

import (
	"context"
	"errors"
	"time"

	"gin-jwt-auth/internal/domain"

	"gorm.io/gorm"
)

// pgUser is the repository-level model for GORM mappings
type pgUser struct {
	ID           uint   `gorm:"primaryKey"`
	FirstName    string `gorm:"size:100;not null"`
	LastName     string `gorm:"size:100"`
	Username     string `gorm:"size:100;uniqueIndex;not null"`
	Password     string `gorm:"not null"`
	Email        string `gorm:"size:255;unique"`
	Phone        string `gorm:"size:15"`
	Role         int    `gorm:"not null"`
	AccessToken  string
	RefreshToken string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// TableName overrides the table name used by User to `users`
func (pgUser) TableName() string {
	return "users"
}

// Convert domain entity to repository entity
func toPGUser(u *domain.User) *pgUser {
	return &pgUser{
		ID:           u.ID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Username:     u.Username,
		Password:     u.Password,
		Email:        u.Email,
		Phone:        u.Phone,
		Role:         int(u.Role),
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

// Convert repository entity to domain entity
func toDomainUser(p *pgUser) *domain.User {
	return &domain.User{
		ID:           p.ID,
		FirstName:    p.FirstName,
		LastName:     p.LastName,
		Username:     p.Username,
		Password:     p.Password,
		Email:        p.Email,
		Phone:        p.Phone,
		Role:         domain.Role(p.Role),
		AccessToken:  p.AccessToken,
		RefreshToken: p.RefreshToken,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new postgres repository for user
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	// AutoMigrate here for ease, though migrations should ideally be separate
	db.AutoMigrate(&pgUser{})
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, u *domain.User) error {
	pgU := toPGUser(u)
	err := r.db.WithContext(ctx).Create(pgU).Error
	if err != nil {
		return err
	}
	u.ID = pgU.ID 
	u.CreatedAt = pgU.CreatedAt
	u.UpdatedAt = pgU.UpdatedAt
	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var p pgUser
	err := r.db.WithContext(ctx).First(&p, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Return nil on not found instead of err for cleaner handling
		}
		return nil, err
	}
	return toDomainUser(&p), nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var p pgUser
	err := r.db.WithContext(ctx).Where("username = ?", username).First(&p).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toDomainUser(&p), nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	var pgUsers []pgUser
	err := r.db.WithContext(ctx).Find(&pgUsers).Error
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for _, p := range pgUsers {
		users = append(users, *toDomainUser(&p))
	}
	return users, nil
}

func (r *userRepository) UpdateTokens(ctx context.Context, id uint, access string, refresh string) error {
	return r.db.WithContext(ctx).Model(&pgUser{}).Where("id = ?", id).Updates(map[string]interface{}{
		"access_token":  access,
		"refresh_token": refresh,
	}).Error
}

func (r *userRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&pgUser{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
