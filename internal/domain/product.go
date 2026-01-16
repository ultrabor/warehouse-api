package domain

import "context"

type Product struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type ProductRepository interface {
	Create(ctx context.Context, p *Product) (int64, error)
	GetByID(ctx context.Context, id int64) (*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]Product, error)
}

type ProductUseCase interface {
	Create(ctx context.Context, p *Product) (int64, error)
	GetByID(ctx context.Context, id int64) (*Product, error)
	Update(ctx context.Context, p *Product) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]Product, error)
}
