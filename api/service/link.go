package service

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"link_shortener/db"
	"link_shortener/db/models"
	"link_shortener/db/repository"
	"link_shortener/utils"
)

type LinkRepo interface {
	FindByUrl(ctx context.Context, url string) (*models.Link, error)
	FindByShort(ctx context.Context, url string) (*models.Link, error)
	Create(ctx context.Context, base, short string) error
	Delete(ctx context.Context, id string) error
}

type Link struct {
	Repo      LinkRepo
	Validator *validator.Validate
}

func (s *Link) Shorten(ctx context.Context, url string) (string, error) {
	short := utils.RandStringBytes(15)
	_, err := s.Repo.FindByShort(ctx, short)
	if err != nil {
		if errors.Is(err, utils.ErrNotFound) {
			return short, nil
		}
		return "", err
	}
	return s.Shorten(ctx, url)
}

func NewLinkService(dbi *db.DB) *Link {
	return &Link{
		Repo:      repository.NewLinkRepo(dbi),
		Validator: validator.New(),
	}
}
