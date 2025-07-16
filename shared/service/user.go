package service

import (
	"crypto/rand"
	"crypto/subtle"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/peteraba/go-frameworks/shared/model"
	"github.com/peteraba/go-frameworks/shared/repo"
	"golang.org/x/crypto/argon2"
)

// The draft RFC recommends[2] time=1, and memory=64*1024 is a sensible number.
// If using that amount of memory (64 MB) is not possible in some contexts then
// the time parameter can be increased to compensate.
const (
	argonTime    = 1
	argonMemory  = 64 * 1024
	argonThreads = 4
	argonKeyLen  = 32
	argonSaltLen = 16
)

var myJWTSigningKey = []byte("GoFr_am3work!s")

const (
	tokenExpiryTime = time.Hour
	tokenIssuer     = "!GOOOOFrrrr3r"
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

	user, err := s.repo.Create(uc, hash, salt)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

var ErrInvalidCredentials = errors.New("invalid credentials")

func (s *UserService) Login(ul model.UserLogin) (string, error) {
	user, err := s.repo.GetByEmail(ul.Email)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token, err: %w", err)
	}

	// Hash the provided password with the stored salt
	hash := argon2.IDKey([]byte(ul.Password), user.PasswordSalt, argonTime, argonMemory, argonThreads, argonKeyLen)

	// Time-attack-resilient comparison of the password against the stored hash
	if subtle.ConstantTimeCompare(hash, user.PasswordHash) != 1 {
		return "", ErrInvalidCredentials
	}

	// Create the Claims
	claims := model.LoggedInUser{
		ID:     user.ID,
		Name:   user.Name,
		Groups: user.Groups,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiryTime)),
			Issuer:    tokenIssuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(myJWTSigningKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign the token, err: %w", err)
	}

	return ss, nil
}

var ErrUnknownClaimsType = errors.New("unknown claims type")

func (s *UserService) TokenToLoggedInUser(tokenString string) (model.LoggedInUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.LoggedInUser{}, func(token *jwt.Token) (any, error) {
		return myJWTSigningKey, nil
	})
	if err != nil {
		log.Fatal(err)
	}

	if claims, ok := token.Claims.(*model.LoggedInUser); ok {
		return *claims, nil
	}

	return model.LoggedInUser{}, ErrUnknownClaimsType
}
