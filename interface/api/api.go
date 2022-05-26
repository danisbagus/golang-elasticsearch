package api

import (
	"github.com/danisbagus/golang-elasticsearch/interface/api/handler"
	"github.com/danisbagus/golang-elasticsearch/repo"
	"github.com/danisbagus/golang-elasticsearch/service"
	"github.com/danisbagus/golang-elasticsearch/utils/config"
	"github.com/labstack/echo"
)

func API(route *echo.Group) {

	es := config.GetESClient()
	productRepo := repo.NewProduct(es)
	productService := service.NewProduct(productRepo)
	productHandler := handler.NewProduct(productService)

	productRoute := route.Group("/product")
	productRoute.POST("", productHandler.Insert)
	productRoute.GET("/:id", productHandler.View)
	productRoute.PUT("/:id", productHandler.Update)
	productRoute.DELETE("/:id", productHandler.Delete)
	productRoute.GET("/search", productHandler.SearchName)

}
