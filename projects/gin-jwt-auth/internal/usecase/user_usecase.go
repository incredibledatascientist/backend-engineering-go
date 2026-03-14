package usecase

import (
	"context"
	"errors"
	"strconv"

	"gin-jwt-auth/internal/domain"
	"gin-jwt-auth/pkg/hash"
	"gin-jwt-auth/pkg/token"
)

var (
	ErrUsernameExists = errors.New("username already exists")
	ErrInvalidLogin   = errors.New("invalid username or password")
	ErrInternalError  = errors.New("internal server error")
	ErrUserNotFound   = errors.New("user not found")
)

type userUsecase struct {
	userRepo   domain.UserRepository
	hasher     hash.PasswordHasher
	tokenMaker token.TokenMaker
}

// NewUserUsecase creates a new usecase layer for user logic
func NewUserUsecase(repo domain.UserRepository, hasher hash.PasswordHasher, tokenMaker token.TokenMaker) domain.UserUsecase {
	return &userUsecase{
		userRepo:   repo,
		hasher:     hasher,
		tokenMaker: tokenMaker,
	}
}

func (u *userUsecase) Signup(ctx context.Context, req domain.UserSignupReq) (domain.AuthResponse, error) {
	exists, err := u.userRepo.CheckUsernameExists(ctx, req.Username)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}
	if exists {
		return domain.AuthResponse{}, ErrUsernameExists
	}

	hashedPassword, err := u.hasher.Hash(req.Password)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}

	user := &domain.User{
		Username: req.Username,
		Password: hashedPassword,
		Role:     domain.RoleUser, // Assign default role
	}

	err = u.userRepo.Create(ctx, user)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}

	return u.generateAuthResponse(ctx, user, "User created successfully")
}

func (u *userUsecase) Login(ctx context.Context, req domain.UserLoginReq) (domain.AuthResponse, error) {
	user, err := u.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}
	if user == nil {
		return domain.AuthResponse{}, ErrInvalidLogin
	}

	isValid := u.hasher.Compare(user.Password, req.Password)
	if !isValid {
		return domain.AuthResponse{}, ErrInvalidLogin
	}

	return u.generateAuthResponse(ctx, user, "Login successful")
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]domain.UserResponse, error) {
	users, err := u.userRepo.GetUsers(ctx)
	if err != nil {
		return nil, ErrInternalError
	}

	var res []domain.UserResponse
	for _, user := range users {
		res = append(res, domain.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		})
	}
	return res, nil
}

func (u *userUsecase) GetUserByID(ctx context.Context, id uint) (*domain.UserResponse, error) {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, ErrInternalError
	}
	if user == nil {
		return nil, ErrUserNotFound
	}
	
	res := &domain.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
	return res, nil
}

// helper to dry up generating tokens and updating DB
func (u *userUsecase) generateAuthResponse(ctx context.Context, user *domain.User, msg string) (domain.AuthResponse, error) {
	idStr := strconv.Itoa(int(user.ID))
	roleStr := strconv.Itoa(int(user.Role))

	accessToken, refreshToken, err := u.tokenMaker.GenerateTokens(idStr, user.Username, roleStr)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}

	err = u.userRepo.UpdateTokens(ctx, user.ID, accessToken, refreshToken)
	if err != nil {
		return domain.AuthResponse{}, ErrInternalError
	}

	return domain.AuthResponse{
		Message: msg,
		User: domain.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
