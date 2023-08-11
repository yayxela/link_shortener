package repository

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"link_shortener/db"
	"link_shortener/db/models"
	"link_shortener/utils"
	"time"
)

type Link struct {
	db *db.DB
}

func (r *Link) FindByUrl(ctx context.Context, url string) (*models.Link, error) {
	model := &models.Link{}
	if err := r.db.Find(ctx, models.LinkCollection, model, bson.M{"base": url}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return model, nil
}

func (r *Link) FindByShort(ctx context.Context, url string) (*models.Link, error) {
	model := &models.Link{}
	if err := r.db.Find(ctx, models.LinkCollection, model, bson.M{"short": url}); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}
	return model, nil
}

func (r *Link) Create(ctx context.Context, base, short string) error {
	model := &models.Link{
		ID:        primitive.NewObjectID(),
		Base:      base,
		Short:     short,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}
	return r.db.Insert(ctx, models.LinkCollection, model)
}

func (r *Link) Delete(ctx context.Context, id string) error {
	return r.db.DeleteById(ctx, models.LinkCollection, id)
}
func NewLinkRepo(dbi *db.DB) *Link {
	return &Link{dbi}
}
