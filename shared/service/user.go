package service

import (
	"crypto/rand"
	"fmt"

	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"golang.org/x/crypto/argon2"
)

const (
	argonTime    = 1
	argonMemory  = 64 * 1024
	argonThreads = 4
	argonKeyLen  = 32
	argonSaltLen = 16
)

type UserService struct {
	repo repo.UserRepo
}

func NewUserService(r repo.UserRepo) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(uc model.UserCreate) (model.User, error) {
	if err := uc.Validate(); err != nil {
		return model.User{}, err
	}

	// Generate salt
	salt := make([]byte, argonSaltLen)
	if _, err := rand.Read(salt); err != nil {
		return model.User{}, fmt.Errorf("failed to generate salt: %w", err)
	}

	// Hash the password with Argon2id
	hash := argon2.IDKey([]byte(uc.Password), salt, argonTime, argonMemory, argonThreads, argonKeyLen)

	// Prepare the user for storage (do not include Password fields)
	userToStore := model.UserCreate{
		Name:   uc.Name,
		Email:  uc.Email,
		Groups: uc.Groups,
	}

	user, err := s.repo.Create(userToStore, salt, hash)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
