package routes

import (
	"github.com/gofiber/fiber/v2"
	"link_shortener/api/handlers"
	"link_shortener/api/service"
)

func RedirectRouter(app fiber.Router, s *service.Link) {
	app.Get(":short", handlers.Redirect(s))
}
