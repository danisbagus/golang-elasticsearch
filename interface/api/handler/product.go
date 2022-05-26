package handler

import (
	"net/http"

	"github.com/danisbagus/golang-elasticsearch/interface/api/handler/product"
	"github.com/danisbagus/golang-elasticsearch/model"
	"github.com/danisbagus/golang-elasticsearch/service"
	"github.com/labstack/echo"
)

type ProductHandler struct {
	service service.IProductService
}

func NewProduct(service service.IProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}

}

func (s *ProductHandler) Insert(c echo.Context) error {
	request := new(product.ProductRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(model.Product)
	product.Name = request.Name
	product.Category = request.Category
	product.Price = request.Price

	res, err := s.service.Insert(c.Request().Context(), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "successfully create product", "data": res})
}

func (s *ProductHandler) Update(c echo.Context) error {
	request := new(product.ProductRequest)
	if err := c.Bind(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	product := new(model.Product)
	product.ID = c.Param("id")
	product.Name = request.Name
	product.Category = request.Category
	product.Price = request.Price

	err := s.service.Update(c.Request().Context(), product)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "successfully update product", "data": product})
}

func (s *ProductHandler) View(c echo.Context) error {
	productID := c.Param("id")
	resProduct, err := s.service.View(c.Request().Context(), productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := new(product.ViewResponse)
	response.ID = resProduct.ID
	response.Name = resProduct.Name
	response.Category = resProduct.Category
	response.Price = resProduct.Price

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "successfully get product", "data": response})
}

func (s *ProductHandler) Delete(c echo.Context) error {
	productID := c.Param("id")
	err := s.service.Delete(c.Request().Context(), productID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "successfully delete product", "data": nil})
}
