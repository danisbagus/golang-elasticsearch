package service

import (
	"context"

	"github.com/danisbagus/golang-elasticsearch/model"
	"github.com/google/uuid"

	"github.com/danisbagus/golang-elasticsearch/repo"
)

type IProductService interface {
	Insert(ctx context.Context, product *model.Product) (*model.Product, error)
}

type ProductService struct {
	repo repo.IProductRepo
}

func NewProduct(repo repo.IProductRepo) IProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Insert(ctx context.Context, product *model.Product) (*model.Product, error) {
	product.ID = uuid.New().String()
	err := s.repo.Insert(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
