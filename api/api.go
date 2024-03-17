package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rk-the-dev/micro-fiber-svc/api/products"
)

func RegisterAPI(app *fiber.App) {

	products.Setup(app)

}
