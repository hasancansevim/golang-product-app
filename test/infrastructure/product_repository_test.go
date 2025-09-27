package infrastructure

import (
	"context"
	"fmt"
	"os"
	"product-app/common/postgresql"
	"product-app/domain"
	"product-app/persistance"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistance.IProductRepository
var dbPool *pgxpool.Pool
var ctx context.Context

func TestMain(m *testing.M) {
	ctx = context.Background()
	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		DbName:                "productapp",
		Username:              "postgres",
		Password:              "123456",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})
	productRepository = persistance.NewProductRepository(dbPool)

	fmt.Println("Before all tests")
	exitCode := m.Run()
	fmt.Println("After all tests")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	TestDataInitialize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateTestData(ctx, dbPool)
}

func TestGetAllProductsByStoreName(t *testing.T) {
	setup(ctx, dbPool)
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
	}
	t.Run("GetAllProductsByStore", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStore("ABC TECH")
		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(ctx, dbPool)

	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "AirFryer",
			Price:    3000.0,
			Discount: 22.0,
			Store:    "ABC TECH",
		},
		{
			Id:       2,
			Name:     "Ütü",
			Price:    1500.0,
			Discount: 10.0,
			Store:    "ABC TECH",
		},
		{
			Id:       3,
			Name:     "Çamaşır Makinesi",
			Price:    10000.0,
			Discount: 15.0,
			Store:    "ABC TECH",
		},
		{
			Id:       4,
			Name:     "Lambader",
			Price:    2000.0,
			Discount: 0.0,
			Store:    "Dekorasyon Sarayı",
		},
	}
	t.Run("GetAllProducts", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestAddProduct(t *testing.T) {
	expectedProducts := []domain.Product{
		{
			Id:       1,
			Name:     "Kupa",
			Price:    100.0,
			Discount: 0.0,
			Store:    "Kırtasiye Merkezi",
		},
	}

	newProduct := domain.Product{
		Name:     "Kupa",
		Price:    100.0,
		Discount: 0.0,
		Store:    "Kırtasiye Merkezi",
	}
	t.Run("AddProduct", func(t *testing.T) {
		productRepository.AddProduct(newProduct)
		actualProducts := productRepository.GetAllProducts()
		assert.Equal(t, 1, len(actualProducts))
		assert.Equal(t, expectedProducts, actualProducts)
	})
	clear(ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(ctx, dbPool)

	t.Run("GetProductById", func(t *testing.T) {
		actualProduct, _ := productRepository.GetById(1)
		_, err := productRepository.GetById(5)

		assert.Equal(t, domain.Product{
			Id: 1, Name: "AirFryer", Price: 3000.0, Discount: 22.0, Store: "ABC TECH",
		}, actualProduct)
		assert.Equal(t, "Product Not Found With Id 5", err.Error())
	})
	clear(ctx, dbPool)
}

//başka servislere giden structure vs gibi şeyleri infrastructure da test ediyoruz

// integration tests
/*
func TestAdd(t *testing.T) {
	t.Run("Test Add", func(t *testing.T) {
		actual := Add(10, 20)
		assert.Equal(t, 30, actual)
	})
}
func Add(x int, y int) int {
	return x + y
}
*/
