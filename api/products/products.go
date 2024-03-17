package products

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rk-the-dev/micro-fiber-svc/models"
	productstore "github.com/rk-the-dev/micro-fiber-svc/stores/productstore"
)

func Setup(app *fiber.App) {
	store := productstore.NewProductStore()
	handler := NewProductHandler(store)
	app.Get("/v1/products", handler.GetAllProducts)
	app.Get("/v1/products/:id", handler.GetProductByID)
	app.Post("/v1/products", handler.Create)
}

type ProductHandler struct {
	store productstore.IProductStore
}

func NewProductHandler(store productstore.IProductStore) *ProductHandler {
	return &ProductHandler{store: store}
}
func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
	Products, _ := h.store.GetAllProducts()

	return c.Status(http.StatusOK).JSON(Products)
}
func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	Products, _ := h.store.GetProductByID(1)
	return c.Status(http.StatusOK).JSON(Products)
}
func (h *ProductHandler) Create(c *fiber.Ctx) error {
	var product models.Product

	// Use BodyParser to parse the request body into the user struct
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	h.store.CreateProduct(product)
	return c.Status(http.StatusOK).JSON(nil)
}
