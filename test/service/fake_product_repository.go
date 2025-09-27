package service

import (
	"product-app/domain"
	"product-app/persistance"
)

type FakeProductRepository struct {
	products []domain.Product
}

func NewFakeProductRepository(initialProducts []domain.Product) persistance.IProductRepository {
	return &FakeProductRepository{
		products: initialProducts,
	}
}

func (fakeProductRepository *FakeProductRepository) GetAllProducts() []domain.Product {
	return fakeProductRepository.products
}

func (fakeProductRepository *FakeProductRepository) GetAllProductsByStore(storeName string) []domain.Product {
	// todo : your turn
	return []domain.Product{}
}

func (fakeProductRepository *FakeProductRepository) AddProduct(product domain.Product) error {
	fakeProductRepository.products = append(fakeProductRepository.products, domain.Product{
		Id:       int64(len(fakeProductRepository.products) + 1),
		Name:     product.Name,
		Price:    product.Price,
		Discount: product.Discount,
		Store:    product.Store,
	})
	return nil
}

func (fakeProductRepository *FakeProductRepository) GetById(productId int64) (domain.Product, error) {
	// todo : your turn
	return domain.Product{}, nil
}
func (fakeProductRepository *FakeProductRepository) DeleteById(productId int64) error {
	// todo : your turn
	return nil
}

func (fakeProductRepository *FakeProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	return nil
}
