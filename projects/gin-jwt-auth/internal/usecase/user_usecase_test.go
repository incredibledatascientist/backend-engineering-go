package usecase_test

import (
	"context"
	"errors"
	"testing"

	"gin-jwt-auth/internal/domain"
	"gin-jwt-auth/internal/usecase"
	"gin-jwt-auth/pkg/hash"
	"gin-jwt-auth/pkg/token"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, u *domain.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	args := m.Called(ctx, id)
	var user *domain.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domain.User)
	}
	return user, args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	args := m.Called(ctx, username)
	var user *domain.User
	if args.Get(0) != nil {
		user = args.Get(0).(*domain.User)
	}
	return user, args.Error(1)
}

func (m *MockUserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *MockUserRepository) UpdateTokens(ctx context.Context, id uint, access string, refresh string) error {
	args := m.Called(ctx, id, access, refresh)
	return args.Error(0)
}

func (m *MockUserRepository) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	args := m.Called(ctx, username)
	return args.Bool(0), args.Error(1)
}

func TestSignup_Success(t *testing.T) {
	mockRepo := new(MockUserRepository)
	hasher := hash.NewBcryptHasher()
	tokenMaker := token.NewJWTMaker("secret")

	userUc := usecase.NewUserUsecase(mockRepo, hasher, tokenMaker)

	req := domain.UserSignupReq{
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("CheckUsernameExists", mock.Anything, req.Username).Return(false, nil)
	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Run(func(args mock.Arguments) {
		u := args.Get(1).(*domain.User)
		u.ID = 1 // simulate DB auto increment
	})
	mockRepo.On("UpdateTokens", mock.Anything, uint(1), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil)

	res, err := userUc.Signup(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, "User created successfully", res.Message)
	assert.Equal(t, "testuser", res.User.Username)
	assert.NotEmpty(t, res.AccessToken)
	assert.NotEmpty(t, res.RefreshToken)

	mockRepo.AssertExpectations(t)
}

func TestSignup_UsernameExists(t *testing.T) {
	mockRepo := new(MockUserRepository)
	hasher := hash.NewBcryptHasher()
	tokenMaker := token.NewJWTMaker("secret")

	userUc := usecase.NewUserUsecase(mockRepo, hasher, tokenMaker)

	req := domain.UserSignupReq{
		Username: "testuser",
		Password: "password123",
	}

	mockRepo.On("CheckUsernameExists", mock.Anything, req.Username).Return(true, nil)

	res, err := userUc.Signup(context.Background(), req)

	assert.Error(t, err)
	assert.True(t, errors.Is(err, usecase.ErrUsernameExists))
	assert.Empty(t, res.AccessToken)

	mockRepo.AssertExpectations(t)
}
