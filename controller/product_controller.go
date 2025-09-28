package controller

import (
	"net/http"
	"product-app/controller/request"
	"product-app/controller/response"
	"product-app/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productService service.IProductService
}

func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (productController *ProductController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/products/:id", productController.GetProductById)
	e.GET("/api/v1/products", productController.GetAllProducts)
	e.POST("/api/v1/products", productController.AddProduct)
	e.PUT("/api/v1/products/:id", productController.UpdatePrice)
	e.DELETE("/api/products/:id", productController.DeleteProductById)
}

func (productController *ProductController) GetProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)

	product, err := productController.productService.GetById(int64(productId))
	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResonse{ErrorDescription: err.Error()})
	}
	return c.JSON(http.StatusOK, response.ToResponse(product))
}

func (productController *ProductController) GetAllProducts(c echo.Context) error {
	param := c.QueryParam("store")
	if len(param) == 0 {
		allProducts := productController.productService.GetAllProducts()
		return c.JSON(http.StatusOK, response.ToResponseList(allProducts))
	}
	productsWithEvenStore := productController.productService.GetAllProductsByStore(param)
	return c.JSON(http.StatusOK, response.ToResponseList(productsWithEvenStore))
}

func (productController *ProductController) AddProduct(c echo.Context) error {
	var addProductRequest request.AddProductRequest
	bindErr := c.Bind(&addProductRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResonse{ErrorDescription: bindErr.Error()})
	}
	err := productController.productService.Add(addProductRequest.ToModel())
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, response.ErrorResonse{ErrorDescription: err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}

func (productController *ProductController) UpdatePrice(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)
	newPrice := c.QueryParam("newPrice")

	if len(newPrice) == 0 {
		return c.JSON(http.StatusBadRequest, response.ErrorResonse{ErrorDescription: "Paramater newPrice is required!"})
	}

	convertedPrice, convertError := strconv.ParseFloat(newPrice, 32)
	if convertError != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResonse{ErrorDescription: "newPrice Format Disrupted!"})
	}
	productController.productService.UpdatePrice(int64(productId), float32(convertedPrice))

	return c.NoContent(http.StatusOK)

}

func (productController *ProductController) DeleteProductById(c echo.Context) error {
	param := c.Param("id")
	productId, _ := strconv.Atoi(param)
	err := productController.productService.DeleteById(int64(productId))

	if err != nil {
		return c.JSON(http.StatusNotFound, response.ErrorResonse{ErrorDescription: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
