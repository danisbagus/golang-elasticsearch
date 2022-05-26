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

type document struct {
	Source interface{} `json:"_source"`
}

type IProductRepo interface {
	Insert(ctx context.Context, product *model.Product) error
	Update(ctx context.Context, product *model.Product) error
	FetchOne(ctx context.Context, id string) (*model.Product, error)
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

func (r *ProductRepo) Update(ctx context.Context, product *model.Product) error {
	body, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("[Update]: marshall data: %w", err)
	}

	req := esapi.UpdateRequest{
		Index:      IndexName,
		DocumentID: product.ID,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{"doc":%s}`, body))),
	}

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res, err := req.Do(ctx, r.es)
	if err != nil {
		return fmt.Errorf("[Update] request: %w", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("[Update]: response: %s", res.String())
	}

	return nil
}

func (r *ProductRepo) FetchOne(ctx context.Context, id string) (*model.Product, error) {
	req := esapi.GetRequest{
		Index:      IndexName,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, TimeOut)
	defer cancel()

	res, err := req.Do(ctx, r.es)
	if err != nil {
		return nil, fmt.Errorf("[FetchOne] request: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("[FetchOne] response: %s", res.String())
	}

	product := new(model.Product)
	var (
		body document
	)
	body.Source = &product

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return nil, fmt.Errorf("[FetchOne] decode: %w", err)
	}

	return product, nil
}
