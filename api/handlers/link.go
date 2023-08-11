package handlers

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"link_shortener/api/entities"
	"link_shortener/api/presenter"
	"link_shortener/api/service"
	"link_shortener/utils"
	"net/http"
)

func AddLink(s *service.Link) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := entities.AddLinkRequest{}
		if err := c.BodyParser(&body); err != nil {
			return presenter.Fail(c, err, http.StatusBadRequest)
		}
		if err := s.Validator.Struct(body); err != nil {
			return presenter.Fail(c, err, http.StatusBadRequest)
		}
		ctx := context.Background()
		link, err := s.Repo.FindByUrl(ctx, body.URL)
		if err != nil && !errors.Is(err, utils.ErrNotFound) {
			return presenter.Fail(c, err, http.StatusInternalServerError)
		}

		if link != nil {
			return presenter.Success(c, entities.AddLinkResponse{ShortURL: link.Short})
		}

		shortUrl, err := s.Shorten(ctx, body.URL)
		if err := s.Repo.Create(context.Background(), body.URL, shortUrl); err != nil {
			return presenter.Fail(c, err, http.StatusInternalServerError)
		}
		return presenter.Success(c, entities.AddLinkResponse{ShortURL: shortUrl})
	}
}

func GetLink(s *service.Link) fiber.Handler {
	return func(c *fiber.Ctx) error {
		short := c.Params("short")
		link, err := s.Repo.FindByShort(context.Background(), short)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return presenter.Fail(c, err, http.StatusNotFound)
			}
			return presenter.Fail(c, err, http.StatusInternalServerError)
		}
		return presenter.Success(c, link)
	}
}

func Redirect(s *service.Link) fiber.Handler {
	return func(c *fiber.Ctx) error {
		short := c.Params("short")
		link, err := s.Repo.FindByShort(context.Background(), short)
		if err != nil {
			if errors.Is(err, utils.ErrNotFound) {
				return presenter.Fail(c, err, http.StatusNotFound)
			}
			return presenter.Fail(c, err, http.StatusInternalServerError)
		}
		return c.Redirect(link.Base, http.StatusMovedPermanently)
	}
}
