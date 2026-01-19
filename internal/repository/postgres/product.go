package postgres

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	domain "github.com/ultrabor/warehouse-api/internal/domain"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pr *ProductRepository) Create(ctx context.Context, p *domain.Product) (int64, error) {
	query := `insert into products (name, description, price, quantity) 
			VALUES ($1, $2, $3, $4) returning id`

	var id int64

	err := pr.db.QueryRowContext(ctx, query, p.Name, p.Description, p.Price, p.Quantity).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("repository creating error: %w", err)
	}

	return id, nil
}

func (pr *ProductRepository) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
	query := `select id, name, description, price, quantity from products where id = $1`

	var product domain.Product
	err := pr.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pr ProductRepository) Update(ctx context.Context, p *domain.Product) error {
	query := `update products set name = $1, description = $2, price = $3, quantity = $4 where id = $5`

	res, err := pr.db.ExecContext(ctx, query, p.Name, p.Description, p.Price, p.Quantity, p.ID)
	if err != nil {
		return fmt.Errorf("repository update err: %w", err)
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("no product was found to update")
	}

	return nil
}

func (pr ProductRepository) Delete(ctx context.Context, id int64) error {
	query := `delete from products where id = $1`

	_, err := pr.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("repository delete error: %w", err)
	}

	return nil
}

func (pr ProductRepository) GetAll(ctx context.Context) ([]domain.Product, error) {
	query := `select id, name, description, price, quantity from products`

	var products []domain.Product

	err := pr.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, fmt.Errorf("repository get all err: %w", err)
	}

	return products, nil
}
