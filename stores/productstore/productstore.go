package productstore

import (
	"fmt"

	"github.com/rk-the-dev/micro-fiber-svc/app"
	"github.com/rk-the-dev/micro-fiber-svc/models"
	"gorm.io/gorm"
)

type IProductStore interface {
	GetAllProducts() ([]models.Product, error)
	GetProductByID(id int) (models.Product, error)
	CreateProduct(product models.Product) error
	UpdateProduct(id int, product models.Product) (models.Product, error)
	DeleteProduct(id int) error
}
type productStore struct {
	db *gorm.DB
}

func NewProductStore() IProductStore {
	db, _ := app.ORMHelper.GetDB(app.DB_SQLLIT)
	// Migrate the schema
	db.AutoMigrate(&models.Product{})
	return &productStore{db: db}
}

func (s *productStore) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	// Find all products.
	result := s.db.Find(&products)

	// Check for errors.
	if result.Error != nil {
		fmt.Println("Error finding products:", result.Error)
		return nil, nil
	}
	return products, nil
}
func (s *productStore) GetProductByID(id int) (models.Product, error) {
	return models.Product{}, nil
}
func (s *productStore) CreateProduct(Product models.Product) error {

	result := s.db.Create(&Product)

	// Check for errors
	if result.Error != nil {
		fmt.Println("Error occurred:", result.Error)
		return result.Error
	}
	return nil
}
func (s *productStore) UpdateProduct(id int, Product models.Product) (models.Product, error) {
	return models.Product{}, nil
}
func (s *productStore) DeleteProduct(id int) error {
	return nil
}
