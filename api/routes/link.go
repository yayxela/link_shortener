package routes

import (
	"github.com/gofiber/fiber/v2"
	"link_shortener/api/handlers"
	"link_shortener/api/service"
)

func LinkRouter(app fiber.Router, s *service.Link) {
	link := app.Group("link")
	link.Get("/:short", handlers.GetLink(s))
	link.Post("/", handlers.AddLink(s))
}
