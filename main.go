package main

import (
	"github.com/abdealt/go_crud/bootstrap"
	"github.com/abdealt/go_crud/repository"
	"github.com/gofiber/fiber/v2"
)

type Repository repository.Repository

func main() {
	app := fiber.New()
	bootstrap.InitApp(app)

}
