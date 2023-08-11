package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"link_shortener/api/routes"
	"link_shortener/api/service"
	"link_shortener/config"
	"link_shortener/db"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	dbi, err := db.New(ctx, &conf)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	linkService := service.NewLinkService(dbi)

	base := app.Group("/")
	routes.RedirectRouter(base, linkService)

	api := app.Group("/api")
	routes.LinkRouter(api, linkService)
	if err = app.Listen(conf.ServerPort); err != nil {
		cancel()
		log.Fatal(err)
	}
}
