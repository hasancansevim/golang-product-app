package main

import (
	"context"
	"product-app/common/app"
	"product-app/common/postgresql"
	"product-app/controller"
	"product-app/persistance"
	"product-app/service"

	"github.com/labstack/echo/v4"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	ctx := context.Background()
	e := echo.New()
	configurationManager := app.NewConfigurationManager()
	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistance.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	producController := controller.NewProductController(productService)

	producController.RegisterRoutes(e)
	e.Start("localhost:8080")
}
