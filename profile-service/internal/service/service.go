package service

import (
	"context"
	"log"
	"profile-service/internal/domain"
	"profile-service/internal/repository"
)

type ProfileService struct {
	Repo *repository.ProfileRepo
}

func (s *ProfileService) GetProfile(
	ctx context.Context,
	email string,
) (*domain.Profile, error) {

	return s.Repo.FindByEmail(ctx, email)
}

func (s *ProfileService) Create(ctx context.Context, user *domain.Profile) error {

	err := s.Repo.Create(ctx, user)
	log.Println(err)
	return err
}
