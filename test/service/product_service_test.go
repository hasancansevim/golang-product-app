package service

import (
	"os"
	"product-app/domain"
	"product-app/service"
	"product-app/service/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

func TestMain(m *testing.M) {
	initialProducts := []domain.Product{
		{
			Id:    1,
			Name:  "AirFryer",
			Price: 1000.0,
			Store: "ABC TECH",
		},
		{
			Id:    2,
			Name:  "Ütü",
			Price: 4000.0,
			Store: "ABC TECH",
		},
	}
	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("ShouldGetAllProducts", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
	})
}

func Test_WhenNoValidationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("WhenNoValidationErrorOccurred_ShouldAddProduct", func(t *testing.T) {
		productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    400.0,
			Discount: 50.0,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü",
			Price:    400.0,
			Discount: 50.0,
			Store:    "ABC TECH",
		}, actualProducts[len(actualProducts)-1])
	})
}

func Test_WhenDiscountIsHigherThen70_ShouldNotAddProduct(t *testing.T) {
	t.Run("WhenDiscountIsHigherThen70_ShouldNotAddProduct", func(t *testing.T) {
		err := productService.Add(model.ProductCreate{
			Name:     "Ütü",
			Price:    400.0,
			Discount: 71.0,
			Store:    "ABC TECH",
		})
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, "Discount cannot be greater than 70", err.Error())
	})
}
