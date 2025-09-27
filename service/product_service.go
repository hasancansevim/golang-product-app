package service

import (
	"errors"
	"product-app/domain"
	"product-app/persistance"
	"product-app/service/model"
)

type IProductService interface {
	Add(productCreate model.ProductCreate) error
	DeleteById(productId int64) error
	GetById(productId int64) (domain.Product, error)
	UpdatePrice(productId int64, newPrice float32) error
	GetAllProducts() []domain.Product
	GetAllProductsByStore(storeName string) []domain.Product
}

type ProductService struct {
	productRepository persistance.IProductRepository
}

func NewProductService(productRepository persistance.IProductRepository) IProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (ProductService *ProductService) Add(productCreate model.ProductCreate) error {
	validateError := validateProductCreate(productCreate)
	if validateError != nil {
		return validateError
	}
	return ProductService.productRepository.AddProduct(domain.Product{
		Name:     productCreate.Name,
		Price:    productCreate.Price,
		Discount: productCreate.Discount,
		Store:    productCreate.Store,
	})
}

func (ProductService *ProductService) DeleteById(productId int64) error {
	return ProductService.productRepository.DeleteById(productId)
}

func (ProductService *ProductService) GetById(productId int64) (domain.Product, error) {
	return ProductService.productRepository.GetById(productId)
}

func (ProductService *ProductService) UpdatePrice(productId int64, newPrice float32) error {
	return ProductService.productRepository.UpdatePrice(productId, newPrice)
}

func (ProductService *ProductService) GetAllProducts() []domain.Product {
	return ProductService.productRepository.GetAllProducts()
}

func (ProductService *ProductService) GetAllProductsByStore(storeName string) []domain.Product {
	return ProductService.GetAllProductsByStore(storeName)
}

func validateProductCreate(productCreate model.ProductCreate) error {
	if productCreate.Discount > 70.0 {
		return errors.New("Discount cannot be greater than 70")
	}
	return nil
}
