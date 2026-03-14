package memory

import (
	"context"
	"sync"
	"time"

	"gin-jwt-auth/internal/domain"
)

type userRepository struct {
	mu     sync.RWMutex
	users  map[uint]*domain.User
	nextID uint
}

// NewUserRepository creates an in-memory repository for user
func NewUserRepository() domain.UserRepository {
	return &userRepository{
		users:  make(map[uint]*domain.User),
		nextID: 1,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = r.nextID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	r.users[r.nextID] = user
	r.nextID++

	return nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, nil // Not found
	}

	return user, nil
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, nil // Not found
}

func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var users []domain.User
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

func (r *userRepository) UpdateTokens(ctx context.Context, id uint, access string, refresh string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, exists := r.users[id]
	if !exists {
		return nil
	}

	user.AccessToken = access
	user.RefreshToken = refresh
	user.UpdatedAt = time.Now()
	return nil
}

func (r *userRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Username == username {
			return true, nil
		}
	}
	return false, nil
}
