package service

import (
	"context"
	"errors"

	"github.com/danisbagus/golang-elasticsearch/model"
	"github.com/google/uuid"

	"github.com/danisbagus/golang-elasticsearch/repo"
)

type IProductService interface {
	Insert(ctx context.Context, product *model.Product) (*model.Product, error)
	Update(ctx context.Context, product *model.Product) error
	View(ctx context.Context, ID string) (*model.Product, error)
	Delete(ctx context.Context, ID string) error
	Search(ctx context.Context, key string, value string) ([]model.Product, error)
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

func (s *ProductService) Update(ctx context.Context, product *model.Product) error {
	err := s.repo.Update(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) View(ctx context.Context, ID string) (*model.Product, error) {
	product, err := s.repo.FetchOne(ctx, ID)
	if err != nil {
		return nil, err
	}

	if product.ID == "" {
		return nil, errors.New("not found")
	}

	return product, nil
}

func (s *ProductService) Delete(ctx context.Context, ID string) error {
	err := s.repo.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ProductService) Search(ctx context.Context, key string, value string) ([]model.Product, error) {
	products, err := s.repo.Search(ctx, key, value)
	if err != nil {
		return nil, err
	}
	return products, nil
}
