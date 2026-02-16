package service

import (
	"auth-service/internal/domain"
	"auth-service/internal/repository"
	"auth-service/pkg/hash"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	secret string
	repo   repository.UserRepository
}

func NewService(s string, r repository.UserRepository) *Service {
	return &Service{
		secret: s,
		repo:   r,
	}
}

func (s *Service) LoginService(
	ctx context.Context,
	email, password string,
) (string, error) {

	user, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		log.Println(err)
		return "", err
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	// compare bcrypt password
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	// JWT with real role from DB
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	return token.SignedString([]byte(s.secret))
}

type SignupResult struct {
	Token string
	User  *domain.User
}

func (s *Service) SignupService(ctx context.Context, role, name, email, password string) (*SignupResult, error) {

	// check if user exists

	exists, err := s.repo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, fmt.Errorf("user already exists")
	}

	// hash password
	hashPass, err := hash.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// create domain user
	user := domain.NewUser(email, role, name, hashPass)

	// save user
	if err := s.repo.Create(ctx, user); err != nil {
		log.Print(err)
		return nil, err
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	signedToken, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	if err := s.syncProfile(user); err != nil {
		return nil, err
	}

	return &SignupResult{
		Token: signedToken,
		User:  user,
	}, nil
}

func (s *Service) syncProfile(user *domain.User) error {

	payload := map[string]string{
		"email": user.Email,
		"name":  user.Name,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post(
		"http://profile:8002/create-profile",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("profile service returned %d", resp.StatusCode)
	}

	return nil
}
