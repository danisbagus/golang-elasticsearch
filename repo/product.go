package repo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/danisbagus/golang-elasticsearch/model"
	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v7"
)

const (
	IndexName = "learn_es_products"
	TimeOut   = time.Second * 15
)

type IProductRepo interface {
	Insert(ctx context.Context, product *model.Product) error
}

type ProductRepo struct {
	es *elasticsearch.Client
}

func NewProduct(es *elasticsearch.Client) IProductRepo {
	return &ProductRepo{
		es: es,
	}
}

func (r *ProductRepo) Insert(ctx context.Context, product *model.Product) error {
	body, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("[Insert]: marshall data: %w", err)
	}

	req := esapi.CreateRequest{
		Index:      IndexName,
		DocumentID: product.ID,
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res, err := req.Do(ctx, r.es)
	if err != nil {
		return fmt.Errorf("[Insert] request: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[Insert]: response: %s", res.String())
	}

	return nil
}
